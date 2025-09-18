package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
	"github.com/sskender/bitgoin/pkg/protocol/messages"
)

type PingHandler struct{}

func (h *PingHandler) Handle(msg protocol.Message, peer protocol.Peer) error {
	log.Printf("handling message '%s' from %s", msg.Command(), peer.Address())
	log.Printf("responding with pong message")

	pingMsg := msg.(*messages.PingMessage)
	pongMsg := messages.NewPongMessage(pingMsg.Nonce)

	return peer.Send(pongMsg)
}
