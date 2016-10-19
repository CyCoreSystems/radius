package radius

import (
	"errors"
	"net"
	"sync"
)

// A Client is the representation of a connection to a RADIUS server
type Client struct {
	opts *Options

	idCounter int64
	idMutex   sync.Mutex

	sessions    []*Session
	sessionLock sync.Mutex
}

// Close closes the client and any related sessions
func (cl *Client) Close() {
	cl.sessionLock.Lock()
	for _, sess := range cl.sessions {
		sess.Stop(NASReboot)
	}

	cl.sessions = make([]*Session, 0) // empty list
	cl.sessionLock.Unlock()
}

// Send sends a request packet to the RADIUS server
func (cl *Client) Send(p *Packet) (*Packet, error) {
	// Inject processing objects
	p.auth = AccountingRequestAuthenticator(cl.opts.SharedSecret)

	// Make connection
	conn, err := net.Dial("udp", cl.opts.Host)
	if err != nil {
		return nil, err
	}

	//TODO: do this another way? atomic int package?
	cl.idMutex.Lock()
	p.ID = Identifier(cl.idCounter)
	cl.idCounter++
	cl.idMutex.Unlock()

	// Write request
	if err := p.Write(conn); err != nil {
		return nil, err
	}

	// Wait for response
	var resp Packet

	if err := resp.Read(conn); err != nil {
		return nil, err
	}

	// verify request and response have matching IDs
	if resp.ID != p.ID {
		return nil, errors.New("Request/Response packets do not have matching IDs")
	}

	return &resp, nil
}
