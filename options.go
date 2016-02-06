package radius

import "golang.org/x/net/context"

// Options defines the options used to connect to a RADIUS server
// (currently only supporting accounting)
type Options struct {
	Host         string
	SharedSecret string

	ChannelBuffer int64
}

// NewClient creates a new client from the given options
func (o *Options) NewClient() (*Client, error) {

	ctx, cancel := context.WithCancel(context.Background())

	cl := &Client{
		opts:     o,
		outgoing: make(chan *Packet, o.ChannelBuffer),
		ctx:      ctx,
		cancel:   cancel,
		packets:  make(map[Identifier]chan *Packet),
	}

	err := cl.Connect()
	return cl, err
}
