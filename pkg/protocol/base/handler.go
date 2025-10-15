package base

import "github.com/sskender/bitgoin/pkg/network"

type Handler interface {
	Handle(msg Message, peer *network.Peer) error
}
