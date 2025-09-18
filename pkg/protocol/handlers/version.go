package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type VersionHandler struct{}

func (h *VersionHandler) Handle(msg protocol.Message) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
