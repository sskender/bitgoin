package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type FeeFilterHandler struct{}

func (h *FeeFilterHandler) Handle(msg protocol.Message) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
