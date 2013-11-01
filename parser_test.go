package parser

import "testing"

func TestEncodePacket(t *testing.T) {
	p := Packet{"message", "test"}
	encoded := encodePacket(p)
	result := "4test"
	if result != encoded {
		t.Errorf("Encoding error: should be \"%s\" got \"%v\".", result, encoded)
	}
}
