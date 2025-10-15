package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol/base"
)

type VersionHandler struct{}

func (h *VersionHandler) Handle(msg base.Message, peer *network.Peer) error {
	log.Printf("handling %s message", msg.Command())
	return nil
}
