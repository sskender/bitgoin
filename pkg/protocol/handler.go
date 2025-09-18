package protocol

type MessageHandler interface {
	Handle(msg Message) error
}
