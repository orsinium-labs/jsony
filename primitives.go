package jsony

import (
	"math"
	"strconv"
)

type (
	Bool       bool
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	UInt       uint
	UInt8      uint8
	UInt16     uint16
	UInt32     uint32
	UInt64     uint64
	UIntPtr    uintptr
	Float32    float32
	Float64    float64
	comment    string
	String     string
	safeString string
	null       bool
)

const Null null = false

var (
	_ Encoder = Bool(true)
	_ Encoder = Int(0)
	_ Encoder = Int8(0)
	_ Encoder = Int16(0)
	_ Encoder = Int32(0)
	_ Encoder = Int64(0)
	_ Encoder = UInt(0)
	_ Encoder = UInt8(0)
	_ Encoder = UInt16(0)
	_ Encoder = UInt32(0)
	_ Encoder = UInt64(0)
	_ Encoder = UIntPtr(0)
	_ Encoder = Float32(0)
	_ Encoder = Float64(0)
	_ Encoder = comment("")
	_ Encoder = String("")
	_ Encoder = SafeString("")
)

func SafeString(s safeString) safeString {
	return s
}

func Comment(s comment) comment {
	return s
}

// EncodeJSON implements [Encoder].
func (v null) EncodeJSON(w *Bytes) {
	w.Extend([]byte("null"))
}

// EncodeJSON implements [Encoder].
func (v Bool) EncodeJSON(w *Bytes) {
	if v {
		w.Extend([]byte("true"))
	} else {
		w.Extend([]byte("false"))
	}
}

// EncodeJSON implements [Encoder].
func (v Int) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v Int8) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v Int16) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v Int32) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v Int64) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UInt) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UInt8) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UInt16) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UInt32) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UInt64) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v UIntPtr) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

// EncodeJSON implements [Encoder].
func (v Float32) EncodeJSON(w *Bytes) {
	abs := float32(math.Abs(float64(v)))
	fmt := byte('f')
	if abs != 0 && (abs < 1e-6 || abs >= 1e21) {
		fmt = 'e'
	}
	w.buf = strconv.AppendFloat(w.buf, float64(v), fmt, -1, 32)
}

// EncodeJSON implements [Encoder].
func (v Float64) EncodeJSON(w *Bytes) {
	abs := math.Abs(float64(v))
	fmt := byte('f')
	if abs != 0 && (abs < 1e-6 || abs >= 1e21) {
		fmt = 'e'
	}
	w.buf = strconv.AppendFloat(w.buf, float64(v), fmt, -1, 64)
}

// EncodeJSON implements [Encoder].
func (v comment) EncodeJSON(w *Bytes) {
	w.Extend([]byte{'/', '/'})
	// we assume that comment has no newlines.
	w.Extend([]byte(v))
	w.Append('\n')
}

// EncodeJSON implements [Encoder].
func (v String) EncodeJSON(w *Bytes) {
	writeString(w, []byte(v))
}

// EncodeJSON implements [Encoder].
func (v safeString) EncodeJSON(w *Bytes) {
	w.Append('"')
	w.Extend([]byte(v))
	w.Append('"')
}
