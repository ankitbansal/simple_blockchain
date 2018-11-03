package simple_blockchain

// https://dev.to/damcosset/trying-to-understand-blockchain-by-making-one-ce4
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type BlockChain struct {
	blocks		[]*Block
}

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
		persistBlockChain()
	} else {
		blockchain.blocks = append(blockchain.blocks, block)
		persistBlock()
	}
}

func getBlockChain() *BlockChain {
	return blockchain_ins
}

func persistBlockChain() {

}

func persistBlock() {

}

func cleanUp() {
	blockchain_ins = nil
}

func loadBlockChain() *BlockChain {
	return blockchain_ins
}

