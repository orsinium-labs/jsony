package jsony

import "unsafe"

// One-letter aliases for people living on the edge.
type (
	B = Bool
	I = Int64
	U = UInt
	S = String
	F = Float64
	A = MixedArray
	O = Object
	M = Map
)

const (
	True  = Bool(true)
	False = Bool(false)
)

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

// Marshal JSON into the end of the given slice of bytes.
//
// Useful if you want to reduce allocations by reusing an existing buffer.
func AppendBytes(buf []byte, e Encoder) []byte {
	b := Bytes{buf}
	e.EncodeJSON(&b)
	return b.buf
}

// Marshal JSON as a string.
func EncodeString(e Encoder) string {
	b := Bytes{}
	e.EncodeJSON(&b)
	return unsafe.String(&b.buf[0], len(b.buf))
}
