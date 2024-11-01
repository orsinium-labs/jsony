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

func BenchmarkInt_Jsony(b *testing.B) {
	v := box(3)
	for i := 0; i < b.N; i++ {
		box(jsony.EncodeBytes(jsony.Int(v)))
	}
}

func BenchmarkInt_Stdlib(b *testing.B) {
	v := box(3)
	for i := 0; i < b.N; i++ {
		b, _ := json.Marshal(v)
		box(b)
	}
}

func BenchmarkObject_Jsony(b *testing.B) {
	name := box("john")
	age := box(42)
	for i := 0; i < b.N; i++ {
		obj := jsony.Object(
			jsony.Field("name", jsony.String(name)),
			jsony.Field("age", jsony.Int(age)),
		)
		box(jsony.EncodeBytes(obj))
	}
}

func BenchmarkObject_Stdlib(b *testing.B) {
	name := box("john")
	age := box(42)
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		v := User{
			Name: name,
			Age:  age,
		}
		_ = json.NewEncoder(&b).Encode(v)
	}
}
