package simple_blockchain

import "testing"

func TestBlockChainShouldInitializeWithGenesisBlock(t *testing.T) {
	bc := createBlockChain()
	genesisBlock := bc.blocks[0]
	if genesisBlock == nil {
		t.Errorf("BlockChain not initialized with genesis block")
	}

}