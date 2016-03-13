package main

import (
	"io"

	"github.com/CyCoreSystems/radius"
)

type acct struct {
	session *radius.Session

	w      io.Writer
	r      io.Reader
	closer io.Closer
}

func (a *acct) Close() error {
	return a.closer.Close()
}

func (a *acct) Write(b []byte) (int, error) {
	c, err := a.w.Write(b)
	if err != nil {
		return c, err
	}

	a.session.AddOutputOctets(int32(c))
	return c, err
}

func (a *acct) Read(b []byte) (int, error) {
	c, err := a.r.Read(b)
	if err != nil {
		return c, err
	}

	a.session.AddInputOctets(int32(c))
	return c, err
}
