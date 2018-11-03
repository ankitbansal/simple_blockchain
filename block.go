package simple_blockchain

import (
	"time"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"fmt"
)

type Supplier struct {
	Id		int
	Name		string
}

type Customer struct {
	Id		int
	Name		string
}

type Rating struct {
	Value 		int
	Timestamp	int64
	Source		string
	Supplier	Supplier
	Customer	Customer
}

type Block struct {
	Records  	[]Rating
	Hash	 	[]byte
	PrevHash 	[]byte
	Timestamp	int64
}

func generateHash(block Block) []byte {
	nonce := 0
	var hash [32]byte;
	hashToString := "default"
	for (hashToString[:2] != "00") {
		timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
		ratingArr := []byte(fmt.Sprintf("%v", block.Records))
		headers := bytes.Join([][]byte{[]byte(strconv.Itoa(nonce)), block.PrevHash, timestamp, ratingArr}, []byte{})
		hash = sha256.Sum256(headers)
		hashToString = base64.URLEncoding.EncodeToString(hash[:])
		nonce++
	}

	fmt.Printf("nonce : %d\n", nonce)
	return hash[:]
}

func genesisBlock() *Block {
	supplier := Supplier{1, "dummySupplier"}
	customer := Customer{1, "dummyConsumer"}
	timestamp := time.Now().Unix()
	rating := Rating{5, timestamp, "dummy", supplier, customer}
	return createBlock([]Rating{rating}, nil)
}

func createBlock(records []Rating, prevHash []byte) *Block {
	var block *Block = &Block{records, nil, prevHash, time.Now().Unix()}
	block.Hash = generateHash(*block)
	return block
}

