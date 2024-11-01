package jsony

type tField struct {
	k String
	v Encoder
}

type tObject []tField

type tArray []Encoder

func Field(k String, v Encoder) tField {
	return tField{k, v}
}

func Object(fs ...tField) Encoder {
	return tObject(fs)
}

func Array(els ...Encoder) Encoder {
	return tArray(els)
}

func (v tObject) EncodeJSON(w *Bytes) {
	next := byte('{')
	for _, f := range v {
		w.Extend([]byte{next, '"'})
		f.k.EncodeJSON(w)
		w.Extend([]byte{'"', ':'})
		f.v.EncodeJSON(w)
		next = ','
	}
	if next == '{' {
		w.Extend([]byte{'{', '}'})
	} else {
		w.Append('}')
	}

}

func (v tArray) EncodeJSON(w *Bytes) {
	if v == nil {
		w.Extend([]byte("null"))
		return
	}
	next := byte('[')
	for _, el := range v {
		w.Append(next)
		el.EncodeJSON(w)
		next = ','
	}
	if next == '[' {
		w.Extend([]byte{'[', ']'})
	} else {
		w.Append(']')
	}

}
