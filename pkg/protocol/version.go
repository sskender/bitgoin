package protocol

import (
	"encoding/binary"
	"encoding/hex"
)

// TODO refactor this shit like envelope

type VersionMessage struct {
	protocolVersion [4]byte
	networkServices [8]byte
	timestamp       [8]byte

	// network address
	networkServicesReceiver [8]byte
	addressReceiver         [16]byte
	portReceiver            [2]byte

	// network address
	networkServicesSender [8]byte
	addressSender         [16]byte
	portSender            [2]byte

	nonce     [8]byte
	userAgent []byte
	height    [4]byte
	relay     [1]byte
}

func NewVersionMessage() *VersionMessage {

	// TODO create new version message with sane defaults

	raw, _ := hex.DecodeString("7f1101000000000000000000ad17835b00000000000000000000000000000000000000000000ffff000000008d20000000000000000000000000000000000000ffff000000008d20f6a8d7a440ec27a11b2f70726f6772616d6d696e67626c6f636b636861696e3a302e312f0000000001")

	m := &VersionMessage{}
	m.Parse(raw)

	return m
}

func (m *VersionMessage) Command() string {
	return "version"
}

func (m *VersionMessage) Parse(raw []byte) error {

	// TODO completely change the parse method

	copy(m.protocolVersion[:], raw[0:4])
	copy(m.networkServices[:], raw[4:12])
	copy(m.timestamp[:], raw[12:20])

	copy(m.networkServicesReceiver[:], raw[20:28])
	copy(m.addressReceiver[:], raw[28:44])
	copy(m.portReceiver[:], raw[44:46])

	copy(m.networkServicesSender[:], raw[46:54])
	copy(m.addressSender[:], raw[54:70])
	copy(m.portSender[:], raw[70:72])

	copy(m.nonce[:], raw[72:80])

	var userAgentSizeRaw [1]byte
	copy(userAgentSizeRaw[:], raw[80:81])

	var userAgentSize uint8 = uint8(userAgentSizeRaw[0])
	m.userAgent = make([]byte, userAgentSize)
	copy(m.userAgent[:], raw[81:81+int(userAgentSize)])

	offset := 81 + int(userAgentSize)
	copy(m.height[:], raw[offset:offset+4])

	copy(m.relay[:], raw[offset+4:offset+4+1])

	return nil
}

func (m *VersionMessage) Serialize() []byte {
	userAgentLen := len(m.userAgent)
	totalLen := 4 + 8 + 8 + // protocolVersion + networkServices + timestamp
		8 + 16 + 2 + // receiver
		8 + 16 + 2 + // sender
		8 + // nonce
		1 + userAgentLen + // user agent varstr (len + data)
		4 + // height
		1 // relay

	raw := make([]byte, totalLen)

	offset := 0

	copy(raw[offset:offset+4], m.protocolVersion[:])
	offset += 4

	copy(raw[offset:offset+8], m.networkServices[:])
	offset += 8

	copy(raw[offset:offset+8], m.timestamp[:])
	offset += 8

	copy(raw[offset:offset+8], m.networkServicesReceiver[:])
	offset += 8

	copy(raw[offset:offset+16], m.addressReceiver[:])
	offset += 16

	copy(raw[offset:offset+2], m.portReceiver[:])
	offset += 2

	copy(raw[offset:offset+8], m.networkServicesSender[:])
	offset += 8

	copy(raw[offset:offset+16], m.addressSender[:])
	offset += 16

	copy(raw[offset:offset+2], m.portSender[:])
	offset += 2

	copy(raw[offset:offset+8], m.nonce[:])
	offset += 8

	// user agent length
	raw[offset] = byte(userAgentLen)
	offset++

	copy(raw[offset:offset+userAgentLen], m.userAgent)
	offset += userAgentLen

	copy(raw[offset:offset+4], m.height[:])
	offset += 4

	copy(raw[offset:offset+1], m.relay[:])
	offset++

	return raw
}

func (m *VersionMessage) ProtocolVersion() uint32 {
	return binary.LittleEndian.Uint32(m.protocolVersion[:])
}

func (m *VersionMessage) PortReceiver() uint16 {
	return binary.LittleEndian.Uint16(m.portReceiver[:])
}

func (m *VersionMessage) PortSender() uint16 {
	return binary.LittleEndian.Uint16(m.portSender[:])
}

func (m *VersionMessage) UserAgent() string {
	return string(m.userAgent)
}

func (m *VersionMessage) StartHeight() uint32 {
	return binary.LittleEndian.Uint32(m.height[:])
}

func (m *VersionMessage) Relay() bool {
	return m.relay[0] != 0
}
