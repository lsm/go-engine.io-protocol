package parser

import (
	"regexp"
	"testing"
)

func TestBasic(t *testing.T) {
	// Encode a simple packet.
	p := Packet{"message", "test"}
	encoded := encodePacket(p)
	result := "4test"
	if result != encoded {
		t.Errorf("Encoding error: should be \"%s\" got \"%v\".", result, encoded)
	}
	// Decode a simple packet.
	decoded := decodePacket(encoded)
	if p != decoded {
		t.Errorf("Decoding error: should be \"%v\" got \"%v\".", p, decoded)
	}
}

func TestEncodingDecoding(t *testing.T) {
	// Allow no data in packet.
	var p Packet
	p.typeStr = "message"
	if p != decodePacket(encodePacket(p)) {
		t.Error("Error: should allow packet without data.")
	}
	// Should encode an open packet.
	pOpen := Packet{"open", "{\"some\":\"json\"}"}
	if pOpen != decodePacket(encodePacket(pOpen)) {
		t.Error("Error: should encode/decode open packet.")
	}

	// Should encode an close packet.
	var pClose Packet
	pClose.typeStr = "close"
	if pClose != decodePacket(encodePacket(pClose)) {
		t.Error("Error: should encode/decode close packet.")
	}

	// Should encode an ping packet.
	pPing := Packet{"ping", "1"}
	if pPing != decodePacket(encodePacket(pPing)) {
		t.Error("Error: should encode/decode ping packet.")
	}

	// Should encode an pong packet.
	pPong := Packet{"pong", "1"}
	if pPong != decodePacket(encodePacket(pPong)) {
		t.Error("Error: should encode/decode pong packet.")
	}

	// Should encode an message packet.
	pMessage := Packet{"message", "some data"}
	if pMessage != decodePacket(encodePacket(pMessage)) {
		t.Error("Error: should encode/decode message packet.")
	}

	// Should encode an upgrade packet.
	pUpgrade := Packet{"upgrade", "some data"}
	if pUpgrade != decodePacket(encodePacket(pUpgrade)) {
		t.Error("Error: should encode/decode upgrade packet.")
	}

	// Should match the encoding format.
	r, _ := regexp.Compile(`^[0-9]$`)
	if true != r.MatchString(encodePacket(p)) {
		t.Error("Error: should match the encoding format.")
	}

	r2, _ := regexp.Compile(`^[0-9]`)
	if true != r2.MatchString(encodePacket(pMessage)) {
		t.Error("Error: should match the encoding format.")
	}

	// Should disallow decoding invaild packet.
	err := Packet{"error", "parser error"}
	if err != decodePacket(":::") {
		t.Error("Error: should disallow bad format.")
	}

	if err != decodePacket("94103") {
		t.Error("Error: should disallow inexistent types.")
	}
}
