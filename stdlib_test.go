package jsony_test

import (
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/jsony"
)

type testMarshaller struct{}

func (testMarshaller) MarshalJSON() ([]byte, error) {
	return []byte("sup"), nil
}

func TestMarshaller(t *testing.T) {
	m := jsony.Marshaller{jsony.Int(13)}
	got, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	if string(got) != "13" {
		t.Fatalf("got: %v", got)
	}
}

func TestFromMarshaller(t *testing.T) {
	m := jsony.FromMarshaller(testMarshaller{})
	got := jsony.EncodeString(m)
	if string(got) != "sup" {
		t.Fatalf("got: %v", got)
	}
}
