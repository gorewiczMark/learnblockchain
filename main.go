package main

import (
	"fmt"

	"github.com/gorewiczMark/learnblockchain/block"
)

func main() {

	chain := block.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block afrer genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("data: %x\n", block.Data)
	}
}
