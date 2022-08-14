package main

import (
	"fmt"

	"github.com/simonieee/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
