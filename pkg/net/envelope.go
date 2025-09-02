package net

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

var NETWORK_MAGIC_MAINNET = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}

type NetworkEnvelope struct {
	networkMagic    [4]byte
	command         [12]byte
	payloadLength   [4]byte
	payloadChecksum [4]byte
	payload         []byte
}

func Parse(raw []byte) (*NetworkEnvelope, error) {
	if len(raw) < 24 {
		return nil, fmt.Errorf("header is too short")
	}

	envelope := &NetworkEnvelope{}

	copy(envelope.networkMagic[:], raw[0:4])

	copy(envelope.command[:], raw[4:16])

	copy(envelope.payloadLength[:], raw[16:20])

	copy(envelope.payloadChecksum[:], raw[20:24])

	payloadLength := envelope.PayloadLength()

	if len(raw) < 24+int(payloadLength) {
		return nil, fmt.Errorf("invalid payload length")
	}

	envelope.payload = make([]byte, payloadLength)
	copy(envelope.payload, raw[24:24+payloadLength])

	if !envelope.verifyChecksum() {
		return nil, fmt.Errorf("invalid checksum")
	}

	return envelope, nil
}

func (e *NetworkEnvelope) Serialize() []byte {
	raw := make([]byte, 24+len(e.payload))

	copy(raw[0:4], e.networkMagic[:])
	copy(raw[4:16], e.command[:])
	copy(raw[16:20], e.payloadLength[:])
	copy(raw[20:24], e.payloadChecksum[:])
	copy(raw[24:], e.payload)

	return raw
}

func (e *NetworkEnvelope) Magic() []byte {
	return e.networkMagic[:]
}

func (e *NetworkEnvelope) Command() string {
	return string(bytes.TrimRight(e.command[:], "\x00"))
}

func (e *NetworkEnvelope) PayloadLength() uint32 {
	return binary.LittleEndian.Uint32(e.payloadLength[:])
}

func (e *NetworkEnvelope) PayloadChecksum() []byte {
	return e.payloadChecksum[:]
}

func (e *NetworkEnvelope) calculateChecksum() []byte {
	hash := sha256.Sum256(e.payload)
	hash = sha256.Sum256(hash[:])
	return hash[:4]
}

func (e *NetworkEnvelope) verifyChecksum() bool {
	calcChecksum := e.calculateChecksum()
	checksum := e.PayloadChecksum()
	return bytes.Equal(calcChecksum, checksum)
}

func (e *NetworkEnvelope) Payload() []byte {
	return e.payload
}
