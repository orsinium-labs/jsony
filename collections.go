package jsony

// A key-value pair, an object's field.
type Field struct {
	K safeString
	V Encoder
}

// An object with fixed keys.
type Object []Field

// An object with dynamically set elements.
//
// The order of items is non-deterministic.
//
// Make sure the keys are either [String] or [SafeString].
// JSON doesn't support non-string keys.
type Map map[Encoder]Encoder

// Array of elements of the same type.
type Array[T Encoder] []T

// Array of [SafeString].
type SafeArray = Array[safeString]

// Array where every element can be of different type.
type MixedArray = Array[Encoder]

// EncodeJSON implements [Encoder].
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

// EncodeJSON implements [Encoder].
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

// EncodeJSON implements [Encoder].
func (v Array[T]) EncodeJSON(w *Bytes) {
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
