package blockchain

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nounce int
}

type BlockChain struct{
	Blocks []*Block
}




func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash ,0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nounce = nonce
	return block

}

func Genesis() *Block {
	return CreateBlock("Genisis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}