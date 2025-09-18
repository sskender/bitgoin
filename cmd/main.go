package main

import (
	"log"

	"github.com/sskender/bitgoin/pkg/node"
)

func main() {
	log.Printf("bitgoin is starting")

	node := node.NewSimpleNode()

	addr := "158.180.239.74:8333"
	node.ConnectPeer(addr)

	node.RunLoop()
}
