package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestPBool(t *testing.T) {
	got := jsony.EncodeString(jsony.PBool(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	x := true
	got = jsony.EncodeString(jsony.PBool(&x))
	want = `true`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPInt(t *testing.T) {
	got := jsony.EncodeString(jsony.PInt(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x int = 13
	got = jsony.EncodeString(jsony.PInt(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPInt8(t *testing.T) {
	got := jsony.EncodeString(jsony.PInt8(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x int8 = 13
	got = jsony.EncodeString(jsony.PInt8(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPInt16(t *testing.T) {
	got := jsony.EncodeString(jsony.PInt16(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x int16 = 13
	got = jsony.EncodeString(jsony.PInt16(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPInt32(t *testing.T) {
	got := jsony.EncodeString(jsony.PInt32(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x int32 = 13
	got = jsony.EncodeString(jsony.PInt32(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPInt64(t *testing.T) {
	got := jsony.EncodeString(jsony.PInt64(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x int64 = 13
	got = jsony.EncodeString(jsony.PInt64(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUInt(t *testing.T) {
	got := jsony.EncodeString(jsony.PUInt(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uint = 13
	got = jsony.EncodeString(jsony.PUInt(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUInt8(t *testing.T) {
	got := jsony.EncodeString(jsony.PUInt8(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uint8 = 13
	got = jsony.EncodeString(jsony.PUInt8(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUInt16(t *testing.T) {
	got := jsony.EncodeString(jsony.PUInt16(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uint16 = 13
	got = jsony.EncodeString(jsony.PUInt16(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUInt32(t *testing.T) {
	got := jsony.EncodeString(jsony.PUInt32(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uint32 = 13
	got = jsony.EncodeString(jsony.PUInt32(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUInt64(t *testing.T) {
	got := jsony.EncodeString(jsony.PUInt64(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uint64 = 13
	got = jsony.EncodeString(jsony.PUInt64(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPUIntPtr(t *testing.T) {
	got := jsony.EncodeString(jsony.PUIntPtr(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x uintptr = 13
	got = jsony.EncodeString(jsony.PUIntPtr(&x))
	want = `13`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPFloat32(t *testing.T) {
	got := jsony.EncodeString(jsony.PFloat32(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x float32 = 13.2
	got = jsony.EncodeString(jsony.PFloat32(&x))
	want = `13.2`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPFloat64(t *testing.T) {
	got := jsony.EncodeString(jsony.PFloat64(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	var x float64 = 13.2
	got = jsony.EncodeString(jsony.PFloat64(&x))
	want = `13.2`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestPString(t *testing.T) {
	got := jsony.EncodeString(jsony.PString(nil))
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}

	x := "sup"
	got = jsony.EncodeString(jsony.PString(&x))
	want = `"sup"`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}
