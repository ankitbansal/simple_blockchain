package simple_blockchain

// https://dev.to/damcosset/trying-to-understand-blockchain-by-making-one-ce4
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

import (
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"log"
	"time"
	"errors"
//	"fmt"
)
type BlockChain struct {
	blocks		[]*Block
}

func createBlockChain() *BlockChain {
	return &BlockChain{[]*Block{}}
}

func addBlock(block *Block) {
	blockchain := loadBlockChain()
	if (blockchain == nil || len(blockchain.blocks) == 0) {
		blockchain = createBlockChain()
		persistBlockChain(blockchain)
	}
	blockchain.blocks = append(blockchain.blocks, block)
	persistBlock(block)
}

func getBlockChain() *BlockChain {
	return loadBlockChain()
}

func persistBlockChain(blockchain *BlockChain) {
	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatal("Error while opening db ")
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("block_bucket"))

		if (bucket != nil) {
			log.Fatal("Bucket shouldn't exist ")
			return errors.New("Bucket shouldn't exist")
		}

		bucket, err := tx.CreateBucket([]byte("block_bucket"))

		if err != nil {
			log.Fatal("Error while persisting bucket ")
			return errors.New("Error while persisting bucket")
		}
		return nil
	})
	defer db.Close()

}

func persistBlock(block *Block) {
	serializedBlock := serializeBlock(block)
	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("block_bucket"))
		err = bucket.Put([]byte(block.Hash), serializedBlock)
		if err != nil {
			log.Fatal(err)
			return errors.New("Error while persisting bucket")
		}
		return nil
	})
	defer db.Close()
}

func serializeBlock(block *Block) []byte {
	var serializedData bytes.Buffer
	encoder := gob.NewEncoder(&serializedData)
	err := encoder.Encode(block)
	if err != nil {
		log.Fatal("encode:", err)
	}
	return serializedData.Bytes()
}

func deserializeBlock(data []byte) *Block {
	var block Block;
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&block)
	return &block
}

func cleanUp() {
	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("block_bucket"))
		if (bucket != nil) {
			err = tx.DeleteBucket([]byte("block_bucket"))
			if err != nil {
				log.Fatal(err)
				return nil
			}
		}
		return nil
	})
	defer db.Close()
}

func loadBlockChain() *BlockChain {
	var blockchain *BlockChain;
	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("block_bucket"))
		blockchain = &BlockChain{[]*Block{}}
		//blocks := []*Block{}
		if (bucket != nil) {
			c := bucket.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				blockchain.blocks = append(blockchain.blocks, deserializeBlock(v))
			}
		}
		return nil
	})
	defer db.Close()
	return blockchain
}

