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

func (v Bool) EncodeJSON(w *SafeWriter) {
	if v {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("true"))
	}
}

func (v Int) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Write(b)
}

func (v Int32) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Write(b)
}

func (v Int64) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendInt(nil, int64(v), 10)
	w.Write(b)
}

func (v UInt) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Write(b)
}

func (v UInt32) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Write(b)
}

func (v UInt64) EncodeJSON(w *SafeWriter) {
	b := strconv.AppendUint(nil, uint64(v), 10)
	w.Write(b)
}

func (v String) EncodeJSON(w *SafeWriter) {
	writeString(w, []byte(v))
}
