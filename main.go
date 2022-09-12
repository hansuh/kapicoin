package main

import (
	"fmt"

	"github.com/hansuh/kapicoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Third Block")
	for _, blocks := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", blocks.Data)
		fmt.Printf("Hash: %s\n", blocks.Hash)
		fmt.Printf("PrevHash: %s\n", blocks.PrevHash)
	}
}
