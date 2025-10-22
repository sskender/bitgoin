package base

import (
	"log"

	"github.com/sskender/bitgoin/pkg/network"
)

type Message interface {
	Command() string
	Parse([]byte) error
	Serialize() []byte
}

func WrapMessage(msg Message) (*network.NetworkEnvelope, error) {
	command := msg.Command()
	payload := msg.Serialize()

	log.Printf("wraping message command '%s' with network envelope", command)

	log.Printf("wrapping payload of len %d", len(payload))

	envelope, err := network.NewNetworkEnvelope(command, payload)
	if err != nil {
		return nil, err
	}

	log.Println("message wrapped with network envelope")

	return envelope, nil
}
