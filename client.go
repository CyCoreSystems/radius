package radius

import (
	"net"
	"sync"

	"golang.org/x/net/context"
)

// A Client is the representation of a connection to a RADIUS server
type Client struct {
	opts     *Options
	conn     net.Conn
	outgoing chan *Packet

	ctx    context.Context
	cancel context.CancelFunc

	lock    sync.Mutex
	packets map[Identifier]chan *Packet
}

// Connect attemps to connect to the RADIUS server
func (cl *Client) Connect() error {
	conn, err := net.Dial("udp", cl.opts.Host)
	if err != nil {
		return err
	}

	cl.conn = conn

	go cl.incomingProcessor()
	go cl.outgoingProcessor()

	return nil
}

// Close closes the client
func (cl *Client) Close() {
	if cl == nil {
		return
	}

	cl.cancel()
}

// Send sends a request packet to the RADIUS server
func (cl *Client) Send(p *Packet) (*Packet, error) {
	// Inject processing objects
	p.auth = AccountingRequestAuthenticator(cl.opts.SharedSecret)

	cl.lock.Lock()
	cl.packets[p.ID] = make(chan *Packet, 1)
	cl.lock.Unlock()

	// Send to outgoing queue
	cl.outgoing <- p

	// wait for response
	cl.lock.Lock()
	x := cl.packets[p.ID]
	cl.lock.Unlock()
	px := <-x

	// clear channel from packet map
	cl.lock.Lock()
	delete(cl.packets, p.ID)
	cl.lock.Unlock()

	return px, nil
}

func (cl *Client) incomingProcessor() {
	for {

		select {
		case <-cl.ctx.Done():
			return
		default:
		}

		var p Packet
		if err := p.Read(cl.conn); err != nil {
			//TODO: log error
			continue
		}

		cl.lock.Lock()
		c, ok := cl.packets[p.ID]
		if !ok {
			// TODO: log mismatching packet
		} else {
			c <- &p
		}
		cl.lock.Unlock()
	}
}

func (cl *Client) outgoingProcessor() {
	defer cl.conn.Close()

	for {
		select {
		case p := <-cl.outgoing:
			if err := p.Write(cl.conn); err != nil {
				//TODO: log error
			}
		case <-cl.ctx.Done():
			return
		}
	}
}
