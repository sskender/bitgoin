package messages

import "github.com/sskender/bitgoin/pkg/protocol"

type PongMessage struct {
	Nonce [8]byte
}

func NewPongMessage(nonce [8]byte) *PongMessage {
	return &PongMessage{Nonce: nonce}
}

func (m *PongMessage) Command() string {
	return protocol.MESSAGE_TYPE_PONG
}

func (m *PongMessage) Parse(raw []byte) error {
	copy(m.Nonce[:], raw[0:8])
	return nil
}

func (m *PongMessage) Serialize() []byte {
	return m.Nonce[:]
}
