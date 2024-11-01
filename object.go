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
