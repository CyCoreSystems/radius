package radius

import "io"

type Reader interface {
	Read(r io.Reader) error
}

type Writer interface {
	Write(w io.Writer) error
}

type WriterFunc func(w io.Writer) error

func (f WriterFunc) Write(w io.Writer) error {
	return f(w)
}
