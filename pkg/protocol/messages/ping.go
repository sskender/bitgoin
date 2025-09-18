package messages

import "github.com/sskender/bitgoin/pkg/protocol"

type PingMessage struct {
	Nonce [8]byte
}

func (m *PingMessage) Command() string {
	return protocol.MESSAGE_TYPE_PING
}

func (m *PingMessage) Parse(raw []byte) error {
	copy(m.Nonce[:], raw[0:8])
	return nil
}

func (m *PingMessage) Serialize() []byte {
	return m.Nonce[:]
}
