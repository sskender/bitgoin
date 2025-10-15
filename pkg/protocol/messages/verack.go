package messages

import (
	"github.com/sskender/bitgoin/pkg/protocol/base"
)

type VerAckMessage struct{}

func NewVerAckMessage() *VerAckMessage {
	return &VerAckMessage{}
}

func (m *VerAckMessage) Command() string {
	return base.MESSAGE_TYPE_VERACK
}

func (m *VerAckMessage) Parse(raw []byte) error {
	return nil
}

func (m *VerAckMessage) Serialize() []byte {
	return []byte{}
}
