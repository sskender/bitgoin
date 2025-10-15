package base

import (
	"fmt"
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

	if len(command) > 12 {
		return nil, fmt.Errorf("command '%s' is invalid - too long", command)
	}

	log.Printf("wrapping payload of len %d", len(payload))

	e := network.NewEmptyNetworkEnvelope()

	e.NetworkMagic = network.NETWORK_MAGIC_MAINNET
	e.Command = command
	e.PayloadLength = uint32(len(payload))
	e.Payload = payload
	e.PayloadChecksum = e.CalculateChecksum()

	log.Println("message wrapped with network envelope")

	return e, nil
}
