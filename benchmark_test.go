package jsony_test

import (
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/jsony"
)

// A black box to ensure that the compiler does not optimize away benchmarked code.
//
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
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	for i := 0; i < b.N; i++ {
		v := User{
			Name: name,
			Age:  age,
		}
		b, _ := json.Marshal(v)
		box(b)
	}
}

func BenchmarkArray_Jsony(b *testing.B) {
	name := box("john")
	age := box(42)
	for i := 0; i < b.N; i++ {
		obj := jsony.Array(jsony.String(name), jsony.Int(age))
		box(jsony.EncodeBytes(obj))
	}
}

func BenchmarkArray_Stdlib(b *testing.B) {
	name := box("john")
	age := box(42)
	for i := 0; i < b.N; i++ {
		v := []any{name, age}
		b, _ := json.Marshal(v)
		box(b)
	}
}
