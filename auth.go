package radius

type Authenticator interface {
	Calculate(p *Packet) []byte
}
