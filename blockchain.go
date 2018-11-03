package simple_blockchain

// https://dev.to/damcosset/trying-to-understand-blockchain-by-making-one-ce4
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

import (
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
)
type BlockChain struct {
	blocks		[]*Block
}

const dbFile = "blockchain.db"
var blockchain_ins *BlockChain;

func createBlockChain(block *Block) *BlockChain {
	blocks := []*Block{block}
	blockchain_ins = &BlockChain{blocks}
	return blockchain_ins
}

func addBlock(block *Block) {
	blockchain := loadBlockChain()
	if (blockchain == nil) {
		blockchain = createBlockChain(block)
		persistBlockChain(blockchain)
	} else {
		blockchain.blocks = append(blockchain.blocks, block)
		persistBlock(block)
	}
}

func getBlockChain() *BlockChain {
	return blockchain_ins
}

func persistBlockChain(blockchain *BlockChain) {
	bolt.Open(dbFile, 0600, nil)
}

func persistBlock(block *Block) {
	serializeBlock(block)
}

func serializeBlock(block *Block) []byte {
	var serializedData bytes.Buffer
	encoder := gob.NewEncoder(&serializedData)
	encoder.Encode(block)
	return serializedData.Bytes()
}

func deserializeBlock(data []byte) *Block {
	var block Block;
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&block)
	return &block
}

func cleanUp() {
	blockchain_ins = nil
}

func loadBlockChain() *BlockChain {
	return blockchain_ins
}

