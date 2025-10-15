package messages

import (
	"encoding/hex"
	"testing"

	"github.com/sskender/bitgoin/pkg/protocol/base"
)

func TestSerializeVerAckMessage(t *testing.T) {
	msg := NewVerAckMessage()

	envelope, err := base.WrapMessage(msg)
	if err != nil {
		t.Fatalf("new envelope could not be created")
	}

	raw := hex.EncodeToString(envelope.Serialize())

	if raw != "f9beb4d976657261636b000000000000000000005df6e0e2" {
		t.Fatalf("invalid serialization")
	}
}
