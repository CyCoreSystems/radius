package radius

import "net"

// A Client is the representation of a connection to a RADIUS server
type Client struct {
	opts *Options
	conn net.Conn
}

// Connect attemps to connect to the RADIUS server
func (cl *Client) Connect() error {
	conn, err := net.Dial("udp", cl.opts.Host)
	if err != nil {
		return err
	}

	cl.conn = conn

	return nil
}

// Close closes the client
func (cl *Client) Close() {
	if cl == nil {
		return
	}

	if cl.conn == nil {
		return
	}

	cl.conn.Close()
}

// Send sends a request packet to the proxy server, waiting for a response packet
func (cl *Client) Send(p *Packet) (*Packet, error) {
	//TODO: this is where we send to a central processing loop that can pair request/responses up by their IDs

	// Inject processing objects
	p.auth = AccountingRequestAuthenticator(cl.opts.SharedSecret)

	return nil, p.Write(cl.conn)
}
