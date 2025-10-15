package protocol

import (
	"log"

	"github.com/sskender/bitgoin/pkg/network"
	"github.com/sskender/bitgoin/pkg/protocol/base"
	"github.com/sskender/bitgoin/pkg/protocol/handlers"
)

var dispatchers = map[string]base.Handler{}

func registerMessageHandler(command string, handler base.Handler) {
	dispatchers[command] = handler
}

func Dispatch(msg base.Message, peer *network.Peer) {
	command := msg.Command()

	log.Printf("dispatching message '%s'", command)

	handler, ok := dispatchers[command]
	if !ok {
		log.Printf("no handler registered for message command '%s'", command)
		return
	}

	err := handler.Handle(msg, peer)
	if err != nil {
		log.Printf("error handling message command '%s': %v", command, err)
	}
}

func init() {
	registerMessageHandler(base.MESSAGE_TYPE_FEEFILTER, &handlers.FeeFilterHandler{})
	registerMessageHandler(base.MESSAGE_TYPE_INV, &handlers.InvHandler{})
	registerMessageHandler(base.MESSAGE_TYPE_PING, &handlers.PingHandler{})
	// TODO pong
	registerMessageHandler(base.MESSAGE_TYPE_SENDCMPCT, &handlers.SendCMPCTHandler{})
	registerMessageHandler(base.MESSAGE_TYPE_VERACK, &handlers.VerAckHandler{})
	registerMessageHandler(base.MESSAGE_TYPE_VERSION, &handlers.VersionHandler{})
}
