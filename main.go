package main

import (
	"fmt"

	"github.com/CodyKat/blockchain_practice/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second")
	chain.AddBlock("Thrid")
	chain.AddBlock("Fourth")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block)
	}
}
