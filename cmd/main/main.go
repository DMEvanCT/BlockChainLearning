package main

import (
	"fmt"
	"github.com/DMEvanCT/BlockChainLearning/blockchain"
	"strconv"
)





func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block to add after Genisis")
	chain.AddBlock("Second Block to add after Genisis")
	chain.AddBlock("Third Block to add after Genisis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previouse hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}