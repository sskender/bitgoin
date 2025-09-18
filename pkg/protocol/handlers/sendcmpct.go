package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type SendCMPCTHandler struct{}

func (h *SendCMPCTHandler) Handle(msg protocol.Message) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
