package main

import (
	"fmt"

	"strconv"

	"github.com/gorewiczMark/learnblockchain/blockchain"

)

func main() {


	chain := blockchain.InitBlockChain()


	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block afrer genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("data: %x\n", block.Data)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
