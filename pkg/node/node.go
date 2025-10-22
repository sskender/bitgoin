package node

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol"
	"github.com/sskender/bitgoin/pkg/protocol/base"
	"github.com/sskender/bitgoin/pkg/protocol/messages"
)

// TODO handle panics

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

	log.Printf("added peer %s", n.peer.Address())

	return nil
}

func (n *Node) handshake() error {
	log.Printf("starting the handshake with peer %s", n.peer.Address())

	versionMsg := messages.NewVersionMessage()

	envelope, err := base.WrapMessage(versionMsg)
	if err != nil {
		return err
	}

	err = n.peer.Send(envelope)
	if err != nil {
		return err
	}

	verackReceived := false
	versionReceived := false

	var msg base.Message

	for {
		if verackReceived && versionReceived {
			break
		}

		envelope, err := n.peer.Receive()
		if err != nil {
			return err
		}

		msg, err = protocol.UnwrapEnvelope(envelope)
		if err != nil {
			return err
		}

		if msg.Command() == base.MESSAGE_TYPE_VERACK {
			log.Printf("got verack message in handshake")

			verackReceived = true
		} else if msg.Command() == base.MESSAGE_TYPE_VERSION {
			log.Printf("got version message in handshake")

			versionReceived = true

			verackMsg := messages.NewVerAckMessage()

			envelope, err := base.WrapMessage(verackMsg)
			if err != nil {
				return err
			}

			err = n.peer.Send(envelope)
			if err != nil {
				return err
			}
		} else {
			log.Printf("ignoring message '%s' in handshake process", msg.Command())
		}

	}

	log.Printf("handshake finished successfully with peer %s", n.peer.Address())

	return nil
}

func (n *Node) RunLoop() error {
	for {
		envelope, err := n.peer.Receive()
		if err != nil {
			log.Printf("error on read from peer %s: %v", n.peer.Address(), err)
			if errors.Is(err, io.EOF) {
				return err
			} else {
				panic(err)
			}
		}

		msg, err := protocol.UnwrapEnvelope(envelope)
		if err != nil {
			panic(err)
		}

		protocol.Dispatch(msg, n.peer)
	}
}

func (n *Node) RunPingLoop() {
	log.Printf("running ping loop")

	// reset RTT of peer on startup
	n.peer.PingRTT = -1

	for {
		time.Sleep(1 * time.Minute)

		pingMsg := messages.NewPingMessage()
		log.Printf("sending ping nonce %x to peer %s", pingMsg.Nonce, n.peer.Address())

		n.peer.PingNonce = pingMsg.Nonce
		n.peer.PingTime = time.Now()

		// TODO handle fails - reconnect peer?

		envelope, err := base.WrapMessage(pingMsg)
		if err != nil {
			panic(err)
		}

		err = n.peer.Send(envelope)
		if err != nil {
			log.Printf("error sending ping to peer %s: %v", n.peer.Address(), err)
			panic(err)
		}
	}
}
