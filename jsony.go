package jsony

import (
	"bytes"
	"io"
)

type Encoder interface {
	EncodeJSON(*SafeWriter)
}

// SafeWriter wraps io.Writer and instead of returning an error stores it.
type SafeWriter struct {
	W   io.Writer
	Err error
}

func (w *SafeWriter) Write(b []byte) {
	_, err := w.W.Write(b)
	if err != nil && w.Err == nil {
		w.Err = err
	}
}

func EncodeString(e Encoder) string {
	iw := bytes.Buffer{}
	sw := SafeWriter{W: &iw}
	e.EncodeJSON(&sw)
	return iw.String()
}

func Write(w io.Writer, e Encoder) error {
	sw := SafeWriter{W: w}
	e.EncodeJSON(&sw)
	return sw.Err
}
