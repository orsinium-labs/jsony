package jsony

import "slices"

// A key-value pair, an [Object]'s field.
type Field struct {
	K safeString
	V Encoder
}

// A key-value pair, an [UnsafeObject]'s field.
type UnsafeField struct {
	K Encoder
	V Encoder
}

// An object with fixed keys.
type Object []Field

// An object with dynamically set elements.
//
// Unlike [Object], kays don't have to be a string literal.
// And unlike [Map], the order of elements is preserved.
//
// Make sure the keys are either [String] or [SafeString].
// JSON doesn't support non-string keys.
type UnsafeObject []UnsafeField

// An object with dynamically set elements.
//
// The order of items is non-deterministic.
// If the order is important, use [UnsafeObject].
//
// Unlike UnsafeObject, Map guarantees that there is only one element
// for each key.
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

// Create a copy of the object with the given fields added to the end.
func (v Object) With(fs ...Field) Object {
	return append(slices.Clone(v), fs...)
}

// Create a copy of the object with the given fields added to the end.
func (v UnsafeObject) With(fs ...UnsafeField) UnsafeObject {
	return append(slices.Clone(v), fs...)
}

// Create a copy of the array with the given elements added to the end.
func (v Array[T]) With(es ...T) Array[T] {
	return append(slices.Clone(v), es...)
}

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
func (v UnsafeObject) EncodeJSON(w *Bytes) {
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
