package simple_blockchain

import (
	"time"
	"bytes"
	"crypto/sha256"
)

// https://dev.to/damcosset/trying-to-understand-blockchain-by-making-one-ce4
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type Supplier struct {
	id		int
	name		string
}

type Customer struct {
	id		int
	name		string
}

type Rating struct {
	value 		int
	timestamp	int64
	source		string
	supplier	Supplier
	customer	Customer
}

type Block struct {
	records  	[]Rating
	hash	 	[]byte
	prevHash 	[]byte
	timestamp	int64
}

type BlockChain struct {
	blocks		[]*Block
}

func main() {

}

func generateHash(block Block) []byte {
	headers := bytes.Join([][]byte{block.prevHash}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

func createBlockChain() *BlockChain {
	block := genesisBlock();
	blocks := []*Block{block}
	return &BlockChain{blocks}
}

func createBlock(records []Rating, prevHash []byte) *Block {
	var block *Block = &Block{records, nil, prevHash, time.Now().Unix()}
	block.hash = []byte("1")
	//block.records =
	return block
}

func genesisBlock() *Block {
	return createBlock(nil, nil)
}
