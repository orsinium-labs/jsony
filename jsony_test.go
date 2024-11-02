package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestEncodeBytes(t *testing.T) {
	got := jsony.EncodeBytes(jsony.Int(13))
	if len(got) != 2 || got[0] != '1' || got[1] != '3' {
		t.Fatalf("got: %v", got)
	}
}

func TestEncodeString(t *testing.T) {
	got := jsony.EncodeString(jsony.Int(13))
	if got != "13" {
		t.Fatalf("got: %v", got)
	}
}
