package protocol

const (
	MESSAGE_TYPE_FEEFILTER = "feefilter"
	MESSAGE_TYPE_INV       = "inv"
	MESSAGE_TYPE_PING      = "ping"
	MESSAGE_TYPE_SENDCMPCT = "sendcmpct"
	MESSAGE_TYPE_VERACK    = "verack"
	MESSAGE_TYPE_VERSION   = "version"
)

type Message interface {
	Command() string
	Parse([]byte) error
	Serialize() []byte
}
