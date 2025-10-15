package messages

import (
	"github.com/sskender/bitgoin/pkg/protocol/base"
)

type InvMessage struct{}

func (m *InvMessage) Command() string {
	return base.MESSAGE_TYPE_INV
}

func (m *InvMessage) Parse(raw []byte) error {
	return nil
}

func (m *InvMessage) Serialize() []byte {
	return []byte{}
}
