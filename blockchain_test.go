package simple_blockchain

import "testing"

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
}

func TestShouldGenerateValidHash(t *testing.T) {
	var block Block = Block{

	}

	hash := generateHash(block);
	if (hash == nil) {
		t.Errorf("Hash can't be empty")
	}

}