package radius

// Options defines the options used to connect to a RADIUS server
// (currently only supporting accounting)
type Options struct {
	Host         string
	SharedSecret string
}

// NewClient creates a new client from the given options
func (o *Options) NewClient() (*Client, error) {

	cl := &Client{
		opts: o,
	}

	return cl, nil
}
