package network

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Peer struct {
	address    string
	connection io.ReadWriteCloser
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func Connect(addr string) (*Peer, error) {
	log.Printf("connecting to peer %s", addr)

	p := Peer{}

	p.address = addr

	conn, err := net.Dial("tcp", p.address)
	if err != nil {
		return nil, err
	}

	p.connection = conn

	p.reader = bufio.NewReader(p.connection)
	p.writer = bufio.NewWriter(p.connection)

	return &p, nil
}

func (p *Peer) Address() string {
	return p.address
}

func (p *Peer) Send(envelope *NetworkEnvelope) error {
	command := envelope.Command

	log.Printf("sending message with command %s to peer %s", command, p.address)

	_, err := p.writer.Write(envelope.Serialize())
	if err != nil {
		return err
	}

	err = p.writer.Flush()
	if err != nil {
		return err
	}

	log.Printf("message with command '%s' sent to peer %s", command, p.address)

	return nil
}

func (p *Peer) Receive() (*NetworkEnvelope, error) {
	log.Printf("reading message from peer %s", p.address)

	envelope := NewEmptyNetworkEnvelope()

	err := envelope.Stream(p.reader)
	if err != nil {
		return nil, err
	}

	command := envelope.Command
	log.Printf("message with command '%s' read from peer %s", command, p.address)

	return envelope, nil
}
