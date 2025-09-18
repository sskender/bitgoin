package messages

type VerAckMessage struct{}

func NewVerAckMessage() *VerAckMessage {
	return &VerAckMessage{}
}

func (m *VerAckMessage) Command() string {
	return "verack"
}

func (m *VerAckMessage) Parse(raw []byte) error {
	return nil
}

func (m *VerAckMessage) Serialize() []byte {
	return []byte{}
}
