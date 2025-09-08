package net

import (
	"fmt"

	"github.com/sskender/bitgoin/pkg/net/messages"
)

type Message interface {
	Command() string
	Parse([]byte) error
	Serialize() []byte
}

var messageRegistry = map[string]Message{
	"version": &messages.VersionMessage{},
	"verack":  &messages.VerAckMessage{},
}

func NewMessage(command string) (Message, error) {
	message, ok := messageRegistry[command]
	if ok {
		return message, nil
	}

	return nil, fmt.Errorf("command '%s' not recognized", command)
}
