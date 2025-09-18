package protocol

type Peer interface {
	Address() string
	Send(msg Message) error
}
