package simple_blockchain

import (
	"testing"
	"encoding/base64"
	"time"
	"fmt"
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

func TestShouldAddBlockInOrder(t *testing.T) {
	bc := createBlockChain()
	supplier := Supplier{1, "dummySupplier"}
	customer := Customer{1, "dummyConsumer"}
	timestamp := time.Now().Unix()
	rating := Rating{5, timestamp, "dummy", supplier, customer}
	newBlock := createBlock([]Rating{rating}, bc.blocks[0].hash)
	addBlock(bc, newBlock)

	if len(bc.blocks) != 2 {
		t.Errorf("Blockchain size should be 2")
	}

	if (bc.blocks[1] != newBlock) {
		t.Errorf("Block is not added properly")
	}

	//this is added to see how overall blockchain looks like
	printBlockChain(bc)
}

func printBlockChain(blockChain *BlockChain) {
	for _, block := range blockChain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.prevHash)
		fmt.Printf("Hash: %x\n", block.hash)
		fmt.Printf("Str Value : %s\n", base64.URLEncoding.EncodeToString(block.hash[:]))
		fmt.Println()
	}
}
