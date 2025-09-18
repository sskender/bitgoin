package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type VerAckHandler struct{}

func (h *VerAckHandler) Handle(msg protocol.Message) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
