package node

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/sskender/bitgoin/pkg/message"
)

// TODO this is not node  - its peer => all the naming is wrong

type Node struct {
	Peer net.Conn
}

func NewNode(addr string) (*Node, error) {

	// TODO this is terrible here

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Node{Peer: conn}, nil
}

func (n *Node) send(msg message.Message) error {
	envelope, err := NewNetworkEnvelopeFromMessage(msg)
	if err != nil {
		return nil
	}

	raw := envelope.Serialize()

	log.Println("sending raw")
	log.Println(raw)

	// TODO get what is written

	_, err = n.Peer.Write(raw)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) read() (message.Message, error) {

	// var raw []byte = []byte{}
	r := bufio.NewReader(n.Peer)

	// TODO implement stream
	buf := make([]byte, 24)
	log.Println("calling read on tcp")
	read, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	fmt.Printf("i have just read header: %d", read)
	fmt.Println(buf)

	envelope, err := Parse(buf)
	if err != nil {
		return nil, err
	}

	fmt.Println(envelope.Magic())

	msg, err := message.NewMessage(envelope.Command())
	if err != nil {
		return nil, err
	}

	err = msg.Parse(envelope.Payload())
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// func (n *Node) waitForHandshake() (message.Message, error) {
// 	for {
// 		msg, err := n.read()
// 		if err != nil {
// 			return nil, err
// 		}

// 		if msg.Command() == "version" || msg.Command() == "verack" {
// 			return msg, nil
// 		}

// 		// TODO handle ping
// 	}
// }

func (n *Node) Handshake() error {

	// TODO handshake with who?

	versionMsg := message.NewVersionMessage()
	verackMsg := message.NewVerAckMessage()

	err := n.send(versionMsg)
	if err != nil {
		return err
	}

	verackReceived := false
	versionReceived := false

	var msg message.Message

	for {
		if verackReceived && versionReceived {
			break
		}

		for {
			log.Println("now reading response")
			msg, err = n.read()
			if err != nil {
				return err
			}

			if msg.Command() == "verack" {
				verackReceived = true
				break
			}

			if msg.Command() == "version" {
				versionReceived = true

				err = n.send(verackMsg)
				if err != nil {
					return err
				}

				break
			}
		}

		// msg, err := n.waitForHandshake()
		// if err != nil {
		// 	return err
		// }

		// if msg.Command() == "verack" { // TODO dont hardcode
		// 	verackReceived = true
		// }

		// if msg.Command() == "version" { // TODO how to messages.VerAckMessage.Command {
		// 	versionReceived = true

		// 	err = n.send(&verackMsg)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
}

// TODO read loop
