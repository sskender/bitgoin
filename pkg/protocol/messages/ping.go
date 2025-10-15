package messages

import (
	"encoding/binary"
	"math/rand/v2"

	"github.com/sskender/bitgoin/pkg/protocol/base"
)

type PingMessage struct {
	Nonce [8]byte
}

func NewPingMessage() *PingMessage {
	nonce := rand.Uint64()

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, nonce)

	return &PingMessage{Nonce: [8]byte(buf)}
}

func (m *PingMessage) Command() string {
	return base.MESSAGE_TYPE_PING
}

func (m *PingMessage) Parse(raw []byte) error {
	copy(m.Nonce[:], raw[0:8])
	return nil
}

func (m *PingMessage) Serialize() []byte {
	return m.Nonce[:]
}
