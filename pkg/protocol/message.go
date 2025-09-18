package protocol

const (
	MESSAGE_TYPE_VERSION = "version"
	MESSAGE_TYPE_VERACK  = "verack"
)

type Message interface {
	Command() string
	Parse([]byte) error
	Serialize() []byte
}
