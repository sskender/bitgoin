package messages

type FeeFilterMessage struct{}

func (m *FeeFilterMessage) Command() string {
	return "feefilter"
}

func (m *FeeFilterMessage) Parse(raw []byte) error {
	return nil
}

func (m *FeeFilterMessage) Serialize() []byte {
	return []byte{}
}
