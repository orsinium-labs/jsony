package jsony_test

import (
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/jsony"
)

// type testMarshaller struct {}

// func (testMarshaller) MarshalJSON() ([]byte, error) {
// 	return []byte("sup"), nil
// }

func TestMarshal(t *testing.T) {
	m := jsony.Marshaller{jsony.Int(13)}
	got, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	if string(got) != "13" {
		t.Fatalf("got: %v", got)
	}
}
