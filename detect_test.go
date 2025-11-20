package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestDetect(t *testing.T) {
	encoder := jsony.Detect(13)
	got := jsony.EncodeString(encoder)
	if got != "13" {
		t.Fatalf("got `%s`, want `13`", got)
	}
}

func TestUnsafeDetect(t *testing.T) {
	cases := []struct {
		given    any
		expected string
	}{
		{13, `13`},
		{-13, `-13`},
		{3.14, `3.14`},
		{true, `true`},
		{"hello", `"hello"`},
		{`say "hi"`, `"say \"hi\""`},
		{uint(13), `13`},
		{uintptr(13), `13`},
		{uint8(13), `13`},
		{uint16(13), `13`},
		{uint32(13), `13`},
		{uint64(13), `13`},
		{int8(13), `13`},
		{int16(13), `13`},
		{int32(13), `13`},
		{int64(13), `13`},
		{float32(3.1415), `3.1415`},
		{float64(3.1415), `3.1415`},
		{nil, `null`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			encoder := jsony.UnsafeDetect(tc.given)
			got := jsony.EncodeString(encoder)
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}
