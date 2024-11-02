package jsony

import "unsafe"

// Encoder is an interface describing objects that can be serialized to JSON.
type Encoder interface {
	EncodeJSON(*Bytes)
}

// Bytes is a slice of bytes that wraps [append] to require no assignment.
type Bytes struct {
	buf []byte
}

func (w *Bytes) Append(b byte) {
	w.buf = append(w.buf, b)
}

func (w *Bytes) Extend(b []byte) {
	w.buf = append(w.buf, b...)
}

// Marshal JSON as a slice of bytes.
func EncodeBytes(e Encoder) []byte {
	b := Bytes{}
	e.EncodeJSON(&b)
	return b.buf
}

// Marshal JSON as a string.
func EncodeString(e Encoder) string {
	b := Bytes{}
	e.EncodeJSON(&b)
	return unsafe.String(&b.buf[0], len(b.buf))
}
