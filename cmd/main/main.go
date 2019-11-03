package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

type BlockChain struct{
	blocks []*Block
}


func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data,  b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block

}

func Genesis() *Block {
	return CreateBlock("Genisis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}





func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block to add after Genisis")
	chain.AddBlock("Second Block to add after Genisis")
	chain.AddBlock("Third Block to add after Genisis")

	for _, block := range chain.blocks {
		fmt.Printf("Previouse hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}