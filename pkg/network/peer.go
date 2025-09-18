package network

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/sskender/bitgoin/pkg/protocol"
)

type Peer struct {
	Address    string
	connection io.ReadWriteCloser
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func Connect(addr string) (*Peer, error) {
	log.Printf("connecting to peer %s", addr)

	p := Peer{}

	p.Address = addr

	conn, err := net.Dial("tcp", p.Address)
	if err != nil {
		return nil, err
	}

	p.connection = conn

	p.reader = bufio.NewReader(p.connection)
	p.writer = bufio.NewWriter(p.connection)

	return &p, nil
}

func (p *Peer) Send(msg protocol.Message) error {
	command := msg.Command()

	log.Printf("sending message with command %s to peer %s", command, p.Address)

	envelope := NewEmptyNetworkEnvelope()
	err := envelope.Wrap(msg)
	if err != nil {
		return err
	}

	_, err = p.writer.Write(envelope.Serialize())
	if err != nil {
		return err
	}

	err = p.writer.Flush()
	if err != nil {
		return err
	}

	log.Printf("message with command '%s' sent to peer %s", command, p.Address)

	return nil
}

func (p *Peer) Read() (protocol.Message, error) {
	log.Printf("reading message from peer %s", p.Address)

	envelope := NewEmptyNetworkEnvelope()

	err := envelope.Stream(p.reader)
	if err != nil {
		return nil, err
	}

	msg, err := envelope.Unwrap()
	if err != nil {
		return nil, err
	}

	log.Printf("message with command '%s' read from peer %s", msg.Command(), p.Address)

	return msg, nil
}
