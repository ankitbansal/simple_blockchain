package simple_blockchain

import (
	"time"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"fmt"
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
	nonce := 0
	var hash [32]byte;
	hashToString := "default"
	for (hashToString[:2] != "00") {
		timestamp := []byte(strconv.FormatInt(block.timestamp, 10))
		ratingArr := []byte(fmt.Sprintf("%v", block.records))
		headers := bytes.Join([][]byte{[]byte(strconv.Itoa(nonce)), block.prevHash, timestamp, ratingArr}, []byte{})
		hash = sha256.Sum256(headers)
		hashToString = base64.URLEncoding.EncodeToString(hash[:])
		nonce++
	}

	fmt.Printf("nonce : %d\n", nonce)
	return hash[:]
}

func createBlockChain() *BlockChain {
	block := genesisBlock();
	blocks := []*Block{block}
	return &BlockChain{blocks}
}

func addBlock(blockchain *BlockChain, block *Block) {
	blockchain.blocks = append(blockchain.blocks, block)
}

func createBlock(records []Rating, prevHash []byte) *Block {
	var block *Block = &Block{records, nil, prevHash, time.Now().Unix()}
	block.hash = generateHash(*block)
	return block
}

func genesisBlock() *Block {
	supplier := Supplier{1, "dummySupplier"}
	customer := Customer{1, "dummyConsumer"}
	timestamp := time.Now().Unix()
	rating := Rating{5, timestamp, "dummy", supplier, customer}
	return createBlock([]Rating{rating}, nil)
}
