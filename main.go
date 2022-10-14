package main

import (
	"fmt"

	"github.com/kensjohnson/kblock/blockchain"
	"github.com/kensjohnson/kblock/transaction"
)

func main() {
	fmt.Println("Hello world")
	blockchain.Hello()

	//create a new Blockchain instance w/mining difficulty of 2
	blockchain := blockchain.CreateBlockshain(2)

	//add Transaction
	t1 := transaction.Transaction{
		SourceSystem: "System 1",
		SourceID:     "ID 1",
		From:         "from me",
		To:           "to you",
		Amount:       15.75,
		Note:         "heres a note",
	}
	blockchain.AddBlock(t1)

	//check for validity
	fmt.Println(blockchain.IsValid())

	//print the entire chain
	for _, block := range blockchain.Chain {
		block.Print()
	}

}
