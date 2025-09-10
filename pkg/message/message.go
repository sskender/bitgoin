package message

import (
	"fmt"
)

type Message interface {
	Command() string
	Parse([]byte) error
	Serialize() []byte
	// Handle() error
}

var messageRegistry = map[string]Message{
	"version": &VersionMessage{},
	"verack":  &VerAckMessage{},
}

func NewMessage(command string) (Message, error) {
	message, ok := messageRegistry[command]
	if ok {
		return message, nil
	}

	return nil, fmt.Errorf("command '%s' not recognized", command)
}

// TODO need to have handler methods in order to make sense at all
