package jsony_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/jsony"
)

//go:noinline
func box[T any](v T) T {
	return v
}

func BenchmarkInt_ToString_Jsony(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		jsony.EncodeString(jsony.Int(v))
	}
}

func BenchmarkInt_ToString_Stdlib(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(v)
	}
}

func BenchmarkInt_ToWriter_Jsony(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}
		_ = jsony.Write(&b, jsony.Int(v))
	}
}

func BenchmarkInt_ToWriter_Stdlib(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}
		_ = json.NewEncoder(&b).Encode(v)
	}
}
