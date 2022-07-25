package main

import (
	"fmt"

	"github.com/kensjohnson/kblock/blockchain"
)

func main() {
	fmt.Println("Hello world")
	blockchain.Hello()

	//create a new Blockchain instance w/mining difficulty of 2
	blockchain := blockchain.CreateBlockshain(2)

	//add transactions
	blockchain.AddBlock("Ken", "phil", 54, "owed phil money")
	blockchain.AddBlock("Phil", "Ken", 23, "partial refund")

	//check for validity
	fmt.Println(blockchain.IsValid())

	//print the entire chain
	for _, block := range blockchain.Chain {
		block.Print()
	}

}
