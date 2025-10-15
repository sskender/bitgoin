package messages

import (
	"github.com/sskender/bitgoin/pkg/protocol/base"
)

type SendCMPCTMessage struct{}

func (m *SendCMPCTMessage) Command() string {
	return base.MESSAGE_TYPE_SENDCMPCT
}

func (m *SendCMPCTMessage) Parse(raw []byte) error {
	return nil
}

func (m *SendCMPCTMessage) Serialize() []byte {
	return []byte{}
}
