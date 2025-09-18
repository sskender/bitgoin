package node

import (
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
	"github.com/sskender/bitgoin/pkg/protocol/handlers"
)

type Dispatcher struct {
	handlers map[string]protocol.MessageHandler
}

func InitializeDispatcher() *Dispatcher {
	d := Dispatcher{
		handlers: make(map[string]protocol.MessageHandler),
	}

	d.Register(protocol.MESSAGE_TYPE_FEEFILTER, &handlers.FeeFilterHandler{})
	d.Register(protocol.MESSAGE_TYPE_INV, &handlers.InvHandler{})
	d.Register(protocol.MESSAGE_TYPE_PING, &handlers.PingHandler{})
	d.Register(protocol.MESSAGE_TYPE_SENDCMPCT, &handlers.SendCMPCTHandler{})
	d.Register(protocol.MESSAGE_TYPE_VERACK, &handlers.VerAckHandler{})
	d.Register(protocol.MESSAGE_TYPE_VERSION, &handlers.VersionHandler{})

	return &d
}

func (d *Dispatcher) Register(cmd string, handler protocol.MessageHandler) {
	d.handlers[cmd] = handler
}

func (d *Dispatcher) Dispatch(msg protocol.Message, peer protocol.Peer) {
	cmd := msg.Command()

	log.Printf("dispatching message '%s'", cmd)

	handler, ok := d.handlers[cmd]
	if !ok {
		log.Printf("no handler registered for message command '%s'", cmd)
		return
	}

	err := handler.Handle(msg, peer)
	if err != nil {
		log.Printf("error handling message command '%s': %v", cmd, err)
	}
}
