package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestObject(t *testing.T) {
	obj := jsony.Object{
		{"name", jsony.String("aragorn")},
		{"age", jsony.Int(82)},
	}
	got := jsony.EncodeString(obj)
	want := `{"name":"aragorn","age":82}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestObject_With(t *testing.T) {
	obj1 := jsony.Object{
		{"name", jsony.String("aragorn")},
		{"age", jsony.Int(82)},
	}
	obj2 := obj1.With(
		jsony.Field{"occupation", jsony.String("strider")},
	)

	{
		got := jsony.EncodeString(obj1)
		want := `{"name":"aragorn","age":82}`
		if got != want {
			t.Fatalf("got `%s`, want `%s`", got, want)
		}
	}
	{
		got := jsony.EncodeString(obj2)
		want := `{"name":"aragorn","age":82,"occupation":"strider"}`
		if got != want {
			t.Fatalf("got `%s`, want `%s`", got, want)
		}
	}
}

func TestUnsafeObject(t *testing.T) {
	obj := jsony.UnsafeObject{
		{jsony.String("name"), jsony.String("aragorn")},
		{jsony.String("age"), jsony.Int(82)},
	}
	got := jsony.EncodeString(obj)
	want := `{"name":"aragorn","age":82}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestUnsafeObject_With(t *testing.T) {
	obj1 := jsony.UnsafeObject{
		{jsony.String("name"), jsony.String("aragorn")},
		{jsony.String("age"), jsony.Int(82)},
	}
	obj2 := obj1.With(
		jsony.UnsafeField{jsony.String("occupation"), jsony.String("strider")},
	)

	{
		got := jsony.EncodeString(obj1)
		want := `{"name":"aragorn","age":82}`
		if got != want {
			t.Fatalf("got `%s`, want `%s`", got, want)
		}
	}
	{
		got := jsony.EncodeString(obj2)
		want := `{"name":"aragorn","age":82,"occupation":"strider"}`
		if got != want {
			t.Fatalf("got `%s`, want `%s`", got, want)
		}
	}
}

func TestObject_Empty(t *testing.T) {
	obj := jsony.Object{}
	got := jsony.EncodeString(obj)
	want := `{}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMap(t *testing.T) {
	obj := jsony.Map{
		jsony.SafeString("name"): jsony.String("aragorn"),
	}
	got := jsony.EncodeString(obj)
	want := `{"name":"aragorn"}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMap_Empty(t *testing.T) {
	obj := jsony.Map{}
	got := jsony.EncodeString(obj)
	want := `{}`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMap_Null(t *testing.T) {
	obj := jsony.Map(nil)
	got := jsony.EncodeString(obj)
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestArray(t *testing.T) {
	obj := jsony.Array[jsony.Int]{13, 7, -5}
	got := jsony.EncodeString(obj)
	want := `[13,7,-5]`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMixedArray(t *testing.T) {
	obj := jsony.MixedArray{jsony.String("aragorn"), jsony.Int(82), jsony.Null}
	got := jsony.EncodeString(obj)
	want := `["aragorn",82,null]`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMixedArray_Empty(t *testing.T) {
	obj := jsony.MixedArray{}
	got := jsony.EncodeString(obj)
	want := `[]`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestMixedArray_Null(t *testing.T) {
	obj := jsony.MixedArray(nil)
	got := jsony.EncodeString(obj)
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}
