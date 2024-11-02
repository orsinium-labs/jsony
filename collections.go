package jsony

type Field struct {
	K safeString
	V Encoder
}

type Object []Field

type Map map[Encoder]Encoder

type Array []Encoder

func (v Object) EncodeJSON(w *Bytes) {
	next := byte('{')
	for _, f := range v {
		w.Append(next)
		f.K.EncodeJSON(w)
		w.Append(':')
		f.V.EncodeJSON(w)
		next = ','
	}
	if next == '{' {
		w.Extend([]byte{'{', '}'})
	} else {
		w.Append('}')
	}
}

func (v Map) EncodeJSON(w *Bytes) {
	if v == nil {
		w.Extend([]byte("null"))
		return
	}
	next := byte('{')
	for k, v := range v {
		w.Append(next)
		k.EncodeJSON(w)
		w.Append(':')
		v.EncodeJSON(w)
		next = ','
	}
	if next == '{' {
		w.Extend([]byte{'{', '}'})
	} else {
		w.Append('}')
	}
}

func (v Array) EncodeJSON(w *Bytes) {
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
