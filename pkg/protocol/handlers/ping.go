package handlers

import (
	"log"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol/base"
	"github.com/sskender/bitgoin/pkg/protocol/messages"
)

type PingHandler struct{}

func (h *PingHandler) Handle(msg base.Message, peer *network.Peer) error {
	log.Printf("handling message '%s' from %s", msg.Command(), peer.Address())
	log.Printf("responding with pong message")

	pingMsg := msg.(*messages.PingMessage)
	pongMsg := messages.NewPongMessage(pingMsg.Nonce)

	envelope, err := base.WrapMessage(pongMsg)
	if err != nil {
		return err
	}

	return peer.Send(envelope)
}
