package jsony

import (
	"unsafe"
)

type Encoder interface {
	EncodeJSON(*Bytes)
}

type Bytes struct {
	buf []byte
}

func (w *Bytes) Append(b byte) {
	w.buf = append(w.buf, b)
}

func (w *Bytes) Extend(b []byte) {
	w.buf = append(w.buf, b...)
}

func EncodeBytes(e Encoder) []byte {
	b := Bytes{}
	e.EncodeJSON(&b)
	return b.buf
}

func EncodeString(e Encoder) string {
	b := Bytes{}
	e.EncodeJSON(&b)
	return unsafe.String(&b.buf[0], len(b.buf))
}
