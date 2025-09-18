package network

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"log"

	"github.com/sskender/bitgoin/pkg/protocol"
)

var NETWORK_MAGIC_MAINNET = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}

type NetworkEnvelope struct {
	NetworkMagic    [4]byte
	Command         string
	PayloadLength   uint32
	PayloadChecksum [4]byte
	Payload         []byte
}

func NewEmptyNetworkEnvelope() *NetworkEnvelope {
	return &NetworkEnvelope{}
}

func (e *NetworkEnvelope) Wrap(msg protocol.Message) error {
	command := msg.Command()
	payload := msg.Serialize()

	log.Printf("wraping message command '%s' with network envelope", command)

	if len(command) > 12 {
		return fmt.Errorf("command '%s' is invalid - too long", command)
	}

	log.Printf("wrapping payload of len %d", len(payload))

	e.NetworkMagic = NETWORK_MAGIC_MAINNET
	e.Command = command
	e.PayloadLength = uint32(len(payload))
	e.Payload = payload
	e.PayloadChecksum = e.calculateChecksum()

	log.Println("message wrapped with network envelope")

	return nil
}

func (e *NetworkEnvelope) Unwrap() (protocol.Message, error) {
	log.Printf("unwrapping network envelope command '%s'", e.Command)

	var msg protocol.Message

	switch e.Command {
	case protocol.MESSAGE_TYPE_VERSION:
		msg = &protocol.VersionMessage{}
	case protocol.MESSAGE_TYPE_VERACK:
		msg = &protocol.VerAckMessage{}
	default:
		return nil, fmt.Errorf("unknown message with command '%s'", e.Command)
	}

	err := msg.Parse(e.Payload)
	if err != nil {
		return nil, err
	}

	log.Printf("unwrapped message command '%s'", msg.Command())

	return msg, nil
}

func (e *NetworkEnvelope) Stream(r io.Reader) error {
	log.Println("reading network envelope header")

	buf := make([]byte, 24)
	read, err := io.ReadFull(r, buf)
	if err != nil {
		return err
	}

	log.Printf("read network envelope header size: %d", read)

	if read < 24 {
		return fmt.Errorf("network envelope header is too short")
	}

	copy(e.NetworkMagic[:], buf[0:4])

	e.Command = string(bytes.TrimRight(buf[4:16], "\x00"))
	e.PayloadLength = binary.LittleEndian.Uint32(buf[16:20])

	copy(e.PayloadChecksum[:], buf[20:24])

	log.Printf("reading network envelope payload with payload length %d", e.PayloadLength)

	e.Payload = make([]byte, e.PayloadLength)
	read, err = io.ReadFull(r, e.Payload)
	if err != nil {
		return err
	}

	log.Printf("read network envelope payload size: %d", read)

	if uint32(read) != e.PayloadLength {
		return fmt.Errorf("network envelope payload is too short")
	}

	if !e.verifyChecksum() {
		return fmt.Errorf("invalid network envelope checksum")
	}

	log.Printf("network envelope stream finished")

	return nil
}

func (e *NetworkEnvelope) Parse(buf []byte) error {
	log.Printf("parsing raw into network envelope len %d", len(buf))

	if len(buf) < 24 {
		return fmt.Errorf("network envelope header is too short")
	}

	copy(e.NetworkMagic[:], buf[0:4])

	e.Command = string(bytes.TrimRight(buf[4:16], "\x00"))
	e.PayloadLength = binary.LittleEndian.Uint32(buf[16:20])

	copy(e.PayloadChecksum[:], buf[20:24])
	copy(e.Payload[:], buf[24:])

	if !e.verifyChecksum() {
		return fmt.Errorf("invalid network envelope checksum")
	}

	log.Printf("network envelope parse finished")

	return nil
}

func (e *NetworkEnvelope) Serialize() []byte {
	log.Println("serializing network envelope")

	buf := make([]byte, 24+e.PayloadLength)

	copy(buf[0:4], e.NetworkMagic[:])
	copy(buf[4:16], []byte(e.Command))

	payloadLenBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(payloadLenBuf[:], e.PayloadLength)
	copy(buf[16:20], payloadLenBuf[:])

	copy(buf[20:24], e.PayloadChecksum[:])
	copy(buf[24:], e.Payload)

	log.Println("serializing network envelope done")

	return buf
}

func (e *NetworkEnvelope) calculateChecksum() [4]byte {
	hash := sha256.Sum256(e.Payload)
	hash = sha256.Sum256(hash[:])

	var hashTrim [4]byte
	copy(hashTrim[:], hash[:4])

	return hashTrim
}

func (e *NetworkEnvelope) verifyChecksum() bool {
	return e.PayloadChecksum == e.calculateChecksum()
}
