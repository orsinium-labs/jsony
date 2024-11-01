package jsony

import (
	"bytes"
	"io"
)

type Encoder interface {
	EncodeJSON(*SafeWriter)
}

// SafeWriter wraps io.Writer.
//
// Buffers writes and stores write errors internally.
type SafeWriter struct {
	w   io.Writer
	buf []byte
	Err error
}

func (w *SafeWriter) WriteByte(b byte) {
	w.buf = append(w.buf, b)
	if len(w.buf) > 1024 {
		w.Flush()
	}
}

func (w *SafeWriter) Write(b []byte) {
	w.buf = append(w.buf, b...)
	if len(w.buf) > 1024 {
		w.Flush()
	}
}

func (w *SafeWriter) Flush() {
	_, err := w.w.Write(w.buf)
	if err != nil && w.Err == nil {
		w.Err = err
	}
	w.buf = w.buf[:0]
}

func EncodeString(e Encoder) string {
	iw := bytes.Buffer{}
	sw := SafeWriter{
		w:   &iw,
		buf: nil,
	}
	e.EncodeJSON(&sw)
	return iw.String()
}

func Write(w io.Writer, e Encoder) error {
	sw := SafeWriter{
		w:   w,
		buf: nil,
	}
	e.EncodeJSON(&sw)
	return sw.Err
}
