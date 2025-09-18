package node

import (
	"errors"
	"io"
	"log"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol"
)

type Node struct {
	peer *network.Peer
}

func NewSimpleNode() *Node {
	return &Node{}
}

func (n *Node) ConnectPeer(addr string) error {
	log.Printf("adding peer %s", addr)

	peer, err := network.Connect(addr)
	if err != nil {
		return err
	}

	n.peer = peer

	err = n.handshake()
	if err != nil {
		return err
	}

	log.Printf("added peer %s", n.peer.Address)

	return nil
}

func (n *Node) handshake() error {
	log.Printf("starting the handshake with peer %s", n.peer.Address)

	versionMsg := protocol.NewVersionMessage()
	err := n.peer.Send(versionMsg)
	if err != nil {
		return err
	}

	verackReceived := false
	versionReceived := false

	var msg protocol.Message

	for {
		if verackReceived && versionReceived {
			break
		}

		msg, err = n.peer.Read()
		if err != nil {
			return err
		}

		if msg.Command() == protocol.MESSAGE_TYPE_VERACK {
			log.Printf("got verack message in handshake")

			verackReceived = true
		} else if msg.Command() == protocol.MESSAGE_TYPE_VERSION {
			log.Printf("got version message in handshake")

			versionReceived = true

			verackMsg := protocol.NewVerAckMessage()
			err = n.peer.Send(verackMsg)
			if err != nil {
				return err
			}
		} else {
			log.Printf("ignoring message '%s' in handshake process", msg.Command())
		}

	}

	log.Printf("handshake finished successfully with peer %s", n.peer.Address)

	return nil
}

func (n *Node) RunLoop() {

	// for each peer { read socket -> dispatch -> maybe write }

	for {
		msg, err := n.peer.Read()
		if err != nil {
			log.Printf("error on read from peer %s: %v", n.peer.Address, err)
			if errors.Is(err, io.EOF) {
				panic(err)
			} else {
				continue
			}
		}

		log.Printf("just got message command '%s'", msg.Command())

		// TODO do something with message
	}
}
