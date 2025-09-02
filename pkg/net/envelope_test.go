package net

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestParseEmptyPayload(t *testing.T) {
	raw, err := hex.DecodeString("f9beb4d976657261636b000000000000000000005df6e0e2")
	if err != nil {
		t.Fatalf("invalid raw message")
	}

	envelope, err := Parse(raw)
	if err != nil {
		t.Fatalf("error parsing envelope: %v", err)
	}

	if !bytes.Equal(envelope.Magic(), NETWORK_MAGIC_MAINNET[:]) {
		t.Fatalf("invalid netowork magic")
	}

	if envelope.Command() != "verack" {
		t.Fatalf("invalid command")
	}

	if envelope.PayloadLength() != 0 {
		t.Fatalf("invalid payload length")
	}

	hash := sha256.Sum256(envelope.Payload())
	hash = sha256.Sum256(hash[:])
	if !bytes.Equal(hash[:4], envelope.PayloadChecksum()[:]) {
		t.Fatalf("invalid payload checksum")
	}

	if !bytes.Equal(envelope.Payload(), []byte{}) {
		t.Fatalf("invalid payload")
	}
}

func TestSerializeEmptyPayload(t *testing.T) {
	t.Fatalf("implement")
}
