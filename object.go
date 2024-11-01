package jsony

type tField struct {
	k String
	v Encoder
}

type tObject []tField

func Field(k String, v Encoder) tField {
	return tField{k, v}
}

func Object(fs ...tField) Encoder {
	return tObject(fs)
}

func (v tObject) EncodeJSON(w *SafeWriter) {
	next := byte('{')
	for _, f := range v {
		w.Write([]byte{next, '"'})
		f.k.EncodeJSON(w)
		w.Write([]byte{'"', ':'})
		f.v.EncodeJSON(w)
		next = ','
	}
	if next == '{' {
		w.Write([]byte{'{', '}'})
	} else {
		w.WriteByte('}')
	}

}
