package handlers

import (
	"log"
	"time"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol/base"
	"github.com/sskender/bitgoin/pkg/protocol/messages"
)

type PongHandler struct{}

func (h *PongHandler) Handle(msg base.Message, peer *network.Peer) error {
	log.Printf("handling message '%s' from %s", msg.Command(), peer.Address())

	pongMsg := msg.(*messages.PongMessage)

	log.Printf("peer ping nonce: %x got pong nonce: %x", peer.PingNonce, pongMsg.Nonce)

	if peer.PingNonce != pongMsg.Nonce {
		log.Printf("ignoring different pong nonce")
		return nil
	}

	// measure round trip time from last sent ping
	peer.PingRTT = time.Since(peer.PingTime)

	log.Printf("peer %s RTT is %dms", peer.Address(), peer.PingRTT.Milliseconds())

	return nil
}
