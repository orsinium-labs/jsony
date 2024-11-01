package jsony

import "strings"

type Writer interface {
	Write([]byte)
}

type Encoder interface {
	EncodeJSON(Writer)
}

type stringWriter struct {
	b strings.Builder
}

func (w *stringWriter) Write(b []byte) {
	_, _ = w.b.Write(b)
}

func EncodeString(e Encoder) string {
	w := stringWriter{}
	e.EncodeJSON(&w)
	return w.b.String()
}
