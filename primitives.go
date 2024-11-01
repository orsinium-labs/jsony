package jsony

import "strconv"

type (
	Bool  bool
	Int   int
	Int32 int32
)

var (
	_ Encoder = Bool(true)
	_ Encoder = Int(0)
	_ Encoder = Int32(0)
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
