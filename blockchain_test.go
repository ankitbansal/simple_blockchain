package simple_blockchain

import (
	"testing"
	"encoding/base64"
	"time"
	"fmt"
	"reflect"
)

func TestBlockChainShouldInitializeWithGenesisBlock(t *testing.T) {
	cleanUp()
	addBlock(genesisBlock())

	bc := getBlockChain()
	genesisBlock := bc.blocks[0]
	if genesisBlock == nil {
		t.Errorf("BlockChain not initialized with genesis block")
	}
	if genesisBlock.PrevHash != nil {
		t.Errorf("Genesis Block must have nil previous hash")
	}
	if genesisBlock.Hash == nil {
		t.Errorf("Genesis Block must have hash value")
	}
	hashToString := base64.URLEncoding.EncodeToString(genesisBlock.Hash[:])
	if hashToString[:2] != "00" {
		t.Errorf("Genesis Block hash is invalid. Expected: %s, got %s", "00", hashToString[:2])
	}
}

func TestShouldAddBlockInOrder(t *testing.T) {
	cleanUp()
	genesisBlock := genesisBlock()
	addBlock(genesisBlock)

	supplier := Supplier{1, "dummySupplier"}
	customer := Customer{1, "dummyConsumer"}
	timestamp := time.Now().Unix()
	rating := Rating{5, timestamp, "dummy", supplier, customer}
	newBlock := createBlock([]Rating{rating}, genesisBlock.Hash)

	addBlock(newBlock)
	bc := getBlockChain()
	if len(bc.blocks) != 2 {
		t.Errorf("Blockchain size should be %d but was %d", 2, len(bc.blocks))
	}

	if (!reflect.DeepEqual(bc.blocks[1], newBlock)) {
		t.Errorf("Block is not added properly")
	}

	//this is added to see how overall blockchain looks like
	printBlockChain(bc)
}

func printBlockChain(blockChain *BlockChain) {
	for _, block := range blockChain.blocks {
		fmt.Printf("Block: %x\n", block)
		fmt.Println()
	}
}

