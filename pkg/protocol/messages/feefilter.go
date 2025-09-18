package messages

import "github.com/sskender/bitgoin/pkg/protocol"

type FeeFilterMessage struct{}

func (m *FeeFilterMessage) Command() string {
	return protocol.MESSAGE_TYPE_FEEFILTER
}

func (m *FeeFilterMessage) Parse(raw []byte) error {
	return nil
}

func (m *FeeFilterMessage) Serialize() []byte {
	return []byte{}
}
