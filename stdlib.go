package jsony

// Marshaller wraps an [Encoder] to make it [json.Marshaller].
//
// Might be useful if you, for some reason, want to mix json with stdlib json.
type Marshaller struct {
	Encoder
}

func (m Marshaller) MarshalJSON() ([]byte, error) {
	b := Bytes{}
	m.EncodeJSON(&b)
	return b.buf, nil
}

// A copy-paste of [json.Marshaller] to not import "encoding/json".
//
// If we have an import from "encoding/json", people will think
// that jsony is just a wrapper.
type marshaller interface {
	MarshalJSON() ([]byte, error)
}

type marshallerWrapper struct {
	m marshaller
}

// FromMarshaller wraps an [json.Marshaller] to make it [Encoder].
func FromMarshaller(m marshaller) Encoder {
	return marshallerWrapper{m}
}

// EncodeJSON implements [Encoder].
func (v marshallerWrapper) EncodeJSON(w *Bytes) {
	b, _ := v.m.MarshalJSON()
	w.Extend(b)
}
