package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type PingHandler struct{}

func (h *PingHandler) Handle(msg protocol.Message) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
