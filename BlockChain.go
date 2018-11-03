package simple_blockchain

import (
	"time"
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
	hash	 	string
	prevHash 	string
	timestamp	int64
}

type BlockChain struct {
	blocks		[]*Block
}

func main() {

}

func createBlockChain() *BlockChain {
	block := genesisBlock();
	blocks := []*Block{block}
	return &BlockChain{blocks}
}

func createBlock(records []Rating, prevHash string) *Block {
	var block *Block = &Block{records, "", prevHash, time.Now().Unix()}
	//block.records =
	return block
}

func genesisBlock() *Block {
	return createBlock(nil, "")
}
