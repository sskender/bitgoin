package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type InvHandler struct{}

func (h *InvHandler) Handle(msg protocol.Message, peer protocol.Peer) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
