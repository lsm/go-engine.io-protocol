// Package parser implements methods for encoding and decoding the engine.io packets.
package parser

import (
	"strconv"
)

// Current protocol version.
var Protocol int = 2

// Packet types.
var Packets map[string]int = map[string]int{
	"open":    0,
	"close":   1,
	"ping":    2,
	"pong":    3,
	"message": 4,
	"upgrade": 5,
	"noop":    6,
}

var packetsList []string = keys(Packets)

// Packet struct.
type Packet struct {
	typeStr string
	data    string
}

// Premade error packet.
var err Packet = Packet{"error", "parser error"}

// Encodes a packet.
//
// 		<packet type id> [ <data> ]
//
// Example:
//		4Hello World
//		2
//		3
func encodePacket(packet Packet) string {
	packetType, isPresent := Packets[packet.typeStr]

	var encoded string

	if isPresent {
		encoded += strconv.Itoa(packetType)
		if len(packet.data) > 0 {
			encoded += packet.data
		}
	}
	return encoded
}

func DecodePacket(str string) Packet {
	var p Packet
	return p
}

func EncodePayload() {

}

func DecodePayload() {

}

func keys(m map[string]int) []string {
	mk := make([]string, len(m))
	i := 0
	for k, _ := range m {
		mk[i] = k
		i++
	}
	return mk
}
