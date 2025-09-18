package messages

type PingMessage struct{}

func (m *PingMessage) Command() string {
	return "ping"
}

func (m *PingMessage) Parse(raw []byte) error {
	return nil
}

func (m *PingMessage) Serialize() []byte {
	return []byte{}
}
