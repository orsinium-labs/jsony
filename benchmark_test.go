package jsony_test

import (
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/jsony"
)

//go:noinline
func box[T any](v T) T {
	return v
}

func BenchmarkInt_Jsony(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		jsony.EncodeString(jsony.Int(v))
	}
}

func BenchmarkInt_Stdlib(b *testing.B) {
	v := box[int](3)
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(v)
	}
}
