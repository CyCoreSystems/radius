package radius

import (
	"encoding/binary"
	"io"
	"sync"
	"time"
)

// A Session represents an Accounting session
type Session struct {
	ID string

	cl *Client

	inputOctets  int32
	outputOctets int32

	attrs []Attribute // both start and stop attributes

	startTime time.Time
	stopTime  time.Time

	once sync.Once
}

func (cl *Client) NewSession(ID string, attrs ...Attribute) *Session {

	sess := &Session{
		ID:    ID,
		cl:    cl,
		attrs: attrs,
	}

	cl.sessionLock.Lock()
	cl.sessions = append(cl.sessions, sess)
	cl.sessionLock.Unlock()

	sess.start(attrs...)

	return sess
}

// Start starts the session
func (s *Session) start(attrs ...Attribute) {

	a := s.attrs
	a = append(a, AccountingStart)
	a = append(a, StringAttribute(AccountingSessionID, s.ID))

	pkt := &Packet{
		Code:       AccountingRequest,
		Attributes: a,
	}

	s.cl.Send(pkt)

	s.startTime = time.Now()
}

// Stop stops the session
func (s *Session) Stop(attrs ...Attribute) {
	s.once.Do(func() {
		s.stopTime = time.Now()

		//sessionTime := s.stopTime.Sub(s.startTime)

		a := s.attrs
		a = append(a, AccountingStop)
		a = append(a, StringAttribute(AccountingSessionID, s.ID))
		a = append(a, Attribute{AccountingOutputOctets, 6, []Writer{WriterFunc(func(w io.Writer) error {
			return binary.Write(w, binary.BigEndian, s.outputOctets)
		})}})
		a = append(a, Attribute{AccountingInputOctets, 6, []Writer{WriterFunc(func(w io.Writer) error {
			return binary.Write(w, binary.BigEndian, s.inputOctets)
		})}})

		for _, ax := range attrs {
			a = append(a, ax)
		}

		/*
			 * TODO: invalid Request Authenticator is raised by freeradius when this is included
			a = append(a, Attribute{AccountingSessionTime, 6, []Writer{WriterFunc(func(w io.Writer) error {
				return binary.Write(w, binary.BigEndian, sessionTime*time.Second)
			})}})
		*/
		pkt := &Packet{
			Code:       AccountingRequest,
			Attributes: a,
		}

		s.cl.Send(pkt)
	})
}

func (s *Session) AddInputOctets(i int32) {
	s.inputOctets += i
}

func (s *Session) AddOutputOctets(i int32) {
	s.outputOctets += i
}
