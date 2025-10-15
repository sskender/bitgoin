package protocol

import (
	"fmt"
	"log"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol/base"
	"github.com/sskender/bitgoin/pkg/protocol/messages"
)

var registry = map[string]base.Message{}

func registerMessage(command string, msg base.Message) {
	registry[command] = msg
}

func UnwrapEnvelope(e *network.NetworkEnvelope) (base.Message, error) {
	command := e.Command

	log.Printf("unwrapping network envelope command '%s'", command)

	msg, ok := registry[command]
	if !ok {
		return nil, fmt.Errorf("unknown message with command '%s'", command)
	}

	err := msg.Parse(e.Payload)
	if err != nil {
		return nil, err
	}

	log.Printf("unwrapped message command '%s'", msg.Command())

	return msg, nil
}

func init() {
	registerMessage(base.MESSAGE_TYPE_FEEFILTER, &messages.FeeFilterMessage{})
	registerMessage(base.MESSAGE_TYPE_INV, &messages.InvMessage{})
	registerMessage(base.MESSAGE_TYPE_PING, &messages.PingMessage{})
	registerMessage(base.MESSAGE_TYPE_PONG, &messages.PongMessage{})
	registerMessage(base.MESSAGE_TYPE_SENDCMPCT, &messages.SendCMPCTMessage{})
	registerMessage(base.MESSAGE_TYPE_VERACK, &messages.VerAckMessage{})
	registerMessage(base.MESSAGE_TYPE_VERSION, &messages.VersionMessage{})
}
