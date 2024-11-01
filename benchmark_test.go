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

func BenchmarkObject_ToWriter_Jsony(b *testing.B) {
	name := box[string]("john")
	age := box[int](42)
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}
		obj := jsony.Object(
			jsony.Field("name", jsony.String(name)),
			jsony.Field("age", jsony.Int(age)),
		)
		_ = jsony.Write(&b, obj)
	}
}

func BenchmarkObject_ToWriter_Stdlib(b *testing.B) {
	name := box[string]("john")
	age := box[int](42)
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
