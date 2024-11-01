package jsony

import (
	"strconv"
)

type (
	Bool   bool
	Int    int
	Int32  int32
	Int64  int64
	UInt   uint
	UInt32 uint32
	UInt64 uint64
	String string
)

var (
	_ Encoder = Bool(true)
	_ Encoder = Int(0)
	_ Encoder = Int32(0)
	_ Encoder = Int64(0)
	_ Encoder = UInt(0)
	_ Encoder = UInt32(0)
	_ Encoder = UInt64(0)
	_ Encoder = String("")
)

func (v Bool) EncodeJSON(w *Bytes) {
	if v {
		w.Extend([]byte("true"))
	} else {
		w.Extend([]byte("true"))
	}
}

func (v Int) EncodeJSON(w *Bytes) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Extend(b)
}

func (v Int32) EncodeJSON(w *Bytes) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Extend(b)
}

func (v Int64) EncodeJSON(w *Bytes) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Extend(b)
}

func (v UInt) EncodeJSON(w *Bytes) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Extend(b)
}

func (v UInt32) EncodeJSON(w *Bytes) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Extend(b)
}

func (v UInt64) EncodeJSON(w *Bytes) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Extend(b)
}

func (v String) EncodeJSON(w *Bytes) {
	writeString(w, []byte(v))
}
