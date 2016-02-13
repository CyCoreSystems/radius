package radius

import "io"

// NOTE: maybe these should be called Readable and Writable?

// Reader defines an object that can be read from the given io.Reader
type Reader interface {
	Read(r io.Reader) error
}

// Writer defines an object that can be written to the given io.Writer
type Writer interface {
	Write(w io.Writer) error
}

// WriterFunc defines a function conversion for the radius.Writer interface
type WriterFunc func(w io.Writer) error

func (f WriterFunc) Write(w io.Writer) error {
	return f(w)
}
