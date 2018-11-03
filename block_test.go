package simple_blockchain

import (
	"testing"
	"encoding/base64"
	"time"
)

func TestShouldGenerateValidHash(t *testing.T) {
	var block Block = Block{
		PrevHash: nil,
		Hash: nil,
		Timestamp: time.Now().Unix(),
	}

	hash := generateHash(block);
	hashToString := base64.URLEncoding.EncodeToString(hash[:])

	if (hash == nil) {
		t.Errorf("Hash can't be empty")
	}
	if (hashToString[:2] != "00") {
		t.Errorf("Hash is invalid. Expected: %s, got %s", "00", hashToString[:2])
	}
}
