package simple_blockchain

import (
	"testing"
	"encoding/base64"
	"time"
)

func TestBlockChainShouldInitializeWithGenesisBlock(t *testing.T) {
	bc := createBlockChain()
	genesisBlock := bc.blocks[0]
	if genesisBlock == nil {
		t.Errorf("BlockChain not initialized with genesis block")
	}
	if genesisBlock.prevHash != nil {
		t.Errorf("Genesis Block must have nil previous hash")
	}
	if genesisBlock.hash == nil {
		t.Errorf("Genesis Block must have hash value")
	}
	hashToString := base64.URLEncoding.EncodeToString(genesisBlock.hash[:])
	if hashToString[:2] != "00" {
		t.Errorf("Genesis Block hash is invalid. Expected: %s, got %s", "00", hashToString[:2])
	}
}

func TestShouldGenerateValidHash(t *testing.T) {
	var block Block = Block{
		prevHash: nil,
		hash: nil,
		timestamp: time.Now().Unix(),
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

