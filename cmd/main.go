package main

import (
	"log"

	"github.com/sskender/bitgoin/pkg/node"
)

func main() {
	log.Printf("bitgoin is starting")

	addr := "158.180.239.74:8333"

	node, err := node.NewNode(addr)
	if err != nil {
		panic(err)
	}

	err = node.Handshake()
	if err != nil {
		panic(err)
	}

	// TODO read in loop
}
