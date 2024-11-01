package jsony

import (
	"math"
	"strconv"
)

type (
	Bool       bool
	Int        int
	Int32      int32
	Int64      int64
	UInt       uint
	UInt32     uint32
	UInt64     uint64
	Float32    float32
	Float64    float64
	String     string
	safeString string
)

var (
	_ Encoder = Bool(true)
	_ Encoder = Int(0)
	_ Encoder = Int32(0)
	_ Encoder = Int64(0)
	_ Encoder = UInt(0)
	_ Encoder = UInt32(0)
	_ Encoder = UInt64(0)
	_ Encoder = Float32(0)
	_ Encoder = Float64(0)
	_ Encoder = String("")
	_ Encoder = SafeString("")
)

func SafeString(s safeString) safeString {
	return s
}

func (v Bool) EncodeJSON(w *Bytes) {
	if v {
		w.Extend([]byte("true"))
	} else {
		w.Extend([]byte("false"))
	}
}

func (v Int) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

func (v Int32) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

func (v Int64) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendInt(w.buf, int64(v), 10)
}

func (v UInt) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

func (v UInt32) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

func (v UInt64) EncodeJSON(w *Bytes) {
	w.buf = strconv.AppendUint(w.buf, uint64(v), 10)
}

func (v Float32) EncodeJSON(w *Bytes) {
	abs := float32(math.Abs(float64(v)))
	fmt := byte('f')
	if abs != 0 && (abs < 1e-6 || abs >= 1e21) {
		fmt = 'e'
	}
	w.buf = strconv.AppendFloat(w.buf, float64(v), fmt, -1, 32)
}

func (v Float64) EncodeJSON(w *Bytes) {
	abs := math.Abs(float64(v))
	fmt := byte('f')
	if abs != 0 && (abs < 1e-6 || abs >= 1e21) {
		fmt = 'e'
	}
	w.buf = strconv.AppendFloat(w.buf, float64(v), fmt, -1, 64)
}

func (v String) EncodeJSON(w *Bytes) {
	writeString(w, []byte(v))
}

func (v safeString) EncodeJSON(w *Bytes) {
	w.Append('"')
	w.Extend([]byte(v))
	w.Append('"')
}
