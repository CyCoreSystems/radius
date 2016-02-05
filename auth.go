package radius

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
)

// An Authenticator is responsible for calculating the Authenticator field of the
// packet
type Authenticator interface {
	Calculate(p *Packet) ([]byte, error)
}

func AccountingRequestAuthenticator(sharedSecret string) Authenticator {
	return &reqAuth{
		sharedSecret: sharedSecret,
	}
}

type reqAuth struct {
	sharedSecret string
}

func (r *reqAuth) Calculate(p *Packet) ([]byte, error) {
	output := bytes.NewBuffer([]byte{})

	p.Code.Write(output)
	p.ID.Write(output)

	binary.Write(output, binary.BigEndian, p.Length())

	for i := 0; i != 16; i++ {
		output.Write([]byte{0})
	}

	for _, a := range p.Attributes {
		if err := a.Write(output); err != nil {
			return nil, err
		}
	}

	output.Write([]byte(r.sharedSecret))

	ret := md5.Sum(output.Bytes())

	return ret[:], nil
}
