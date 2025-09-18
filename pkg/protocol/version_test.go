package protocol

import (
	"encoding/hex"
	"testing"
)

func TestParseVersionMessage(t *testing.T) {
	raw, err := hex.DecodeString("7f1101000000000000000000ad17835b00000000000000000000000000000000000000000000ffff000000008d20000000000000000000000000000000000000ffff000000008d20f6a8d7a440ec27a11b2f70726f6772616d6d696e67626c6f636b636861696e3a302e312f0000000001")
	if err != nil {
		t.Fatalf("invalid raw version message")
	}

	message := VersionMessage{}

	err = message.Parse(raw)
	if err != nil {
		t.Fatalf("error parsing version message: %v", err)
	}

	if message.ProtocolVersion() != 70015 {
		t.Fatalf("invalid protocol version")
	}

	if message.PortReceiver() != 8333 {
		t.Fatalf("invalid port receiver")
	}

	if message.PortSender() != 8333 {
		t.Fatalf("invalid port sender")
	}

	if message.UserAgent() != "/programmingblockchain:0.1/" {
		t.Fatalf("invalid user agent")
	}

	if message.StartHeight() != 0 {
		t.Fatalf("invalid start height")
	}

	if !message.Relay() {
		t.Fatalf("invalid relay flag")
	}
}
