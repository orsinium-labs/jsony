package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestObject(t *testing.T) {
	obj := jsony.Object(
		jsony.Field("name", jsony.String("aragorn")),
		jsony.Field("age", jsony.Int(82)),
	)
	got := jsony.EncodeString(obj)
	want := `{"name":"aragorn","age":82}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestObject_Empty(t *testing.T) {
	obj := jsony.Object()
	got := jsony.EncodeString(obj)
	want := `{}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestArray(t *testing.T) {
	obj := jsony.Array(jsony.String("aragorn"), jsony.Int(82), jsony.Null)
	got := jsony.EncodeString(obj)
	want := `["aragorn",82,null]`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestArray_Empty(t *testing.T) {
	obj := jsony.Array()
	got := jsony.EncodeString(obj)
	want := `[]`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}
