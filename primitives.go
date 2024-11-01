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

func (v Bool) EncodeJSON(w Writer) {
	if v {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("true"))
	}
}

func (v Int) EncodeJSON(w Writer) {
	w.Write([]byte(strconv.Itoa(int(v))))
}

func (v Int32) EncodeJSON(w Writer) {
	w.Write([]byte(strconv.FormatInt(int64(v), 10)))
}
