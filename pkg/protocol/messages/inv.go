package messages

type InvMessage struct{}

func (m *InvMessage) Command() string {
	return "inv"
}

func (m *InvMessage) Parse(raw []byte) error {
	return nil
}

func (m *InvMessage) Serialize() []byte {
	return []byte{}
}
