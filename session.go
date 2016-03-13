package radius

import (
	"sync"
	"time"
)

// A Session represents an Accounting session
type Session struct {
	ID string

	cl *Client

	inputOctets  int64
	outputOctets int64

	attrs []Attribute // both start and stop attributes

	startTime time.Time
	stopTIme  time.Time

	once sync.Once
}

func (cl *Client) NewSession(ID string, attrs ...Attribute) *Session {

	sess := &Session{
		ID:    ID,
		cl:    cl,
		attrs: attrs,
	}

	sess.start(attrs...)

	return sess
}

// Start starts the session
func (s *Session) start(attrs ...Attribute) {
	//TODO: Send start packet
	//TODO: set startTime
}

// Stop stops the session
func (s *Session) Stop(terminateCause Attribute) {
	s.once.Do(func() {
		//TODO: set end time
		//TODO: calculate session time
		//TODO: send stop packet
	})
}

func (s *Session) AddInputOctets(i int64) {
	s.inputOctets += i
}

func (s *Session) AddOutputOctets(i int64) {
	s.outputOctets += i
}
