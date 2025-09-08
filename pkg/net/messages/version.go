package messages

import "encoding/binary"

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

func NewVersionMessage(raw []byte) (*VersionMessage, error) {
	m := VersionMessage{}

	err := m.Parse(raw)
	if err != nil {
		return nil, err
	}

	return &m, nil
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
	return []byte{}
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
