package simple_blockchain

// https://dev.to/damcosset/trying-to-understand-blockchain-by-making-one-ce4
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type BlockChain struct {
	blocks		[]*Block
}

var blockchain_ins *BlockChain;

func createBlockChain() *BlockChain {
	block := genesisBlock();
	blocks := []*Block{block}
	persistBlockChain()
	blockchain_ins = &BlockChain{blocks}
	return blockchain_ins
}

func addBlock(block *Block) {
	blockchain := loadBlockChain()
	if (blockchain == nil) {
		blockchain = createBlockChain()
	}

	blockchain.blocks = append(blockchain.blocks, block)
	persistBlock()
}

func persistBlockChain() {

}

func persistBlock() {

}

func loadBlockChain() *BlockChain {
	return blockchain_ins
}

