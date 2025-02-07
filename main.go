package main

import (
	"github.com/CodyKat/blockchain_practice/explorer"
	"github.com/CodyKat/blockchain_practice/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
