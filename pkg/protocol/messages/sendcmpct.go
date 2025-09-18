package messages

type SendCMPCTMessage struct{}

func (m *SendCMPCTMessage) Command() string {
	return "sendcmpct"
}

func (m *SendCMPCTMessage) Parse(raw []byte) error {
	return nil
}

func (m *SendCMPCTMessage) Serialize() []byte {
	return []byte{}
}
