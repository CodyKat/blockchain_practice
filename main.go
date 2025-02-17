package main

import (
	"github.com/CodyKat/blockchain_practice/blockchain"
	"github.com/CodyKat/blockchain_practice/cli"
	"github.com/CodyKat/blockchain_practice/db"
)

func main() {
	defer db.Close()
	blockchain.BlockChain()
	cli.Start()
}
