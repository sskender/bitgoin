package protocol

type MessageHandler interface {
	Handle(msg Message, peer Peer) error
}
