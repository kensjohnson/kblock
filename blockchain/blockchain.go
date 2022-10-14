// Package Blockchain contains the structures and functions needed to
// support building and using a Blockchain
package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kensjohnson/kblock/transaction"
)

// Block is the core structure and is managed ont he BlockChain
type Block struct {
	// TODO #8 convert to a Transaction structure persisted as JSON
	Data         string
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	POW          int
	BlockNote    string
}

// calculateHash is an PKG visible function to derive the block hash for a block of data
func (b Block) calculateHash() string {
	//	data, _ := json.Marshal(b.Data)
	//	fmt.Printf("DEBUG: Marshalled data: <%s>\n", data)
	blockData := b.PreviousHash + b.Data + b.Timestamp.String() + strconv.Itoa(b.POW)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// mine is a PKG visible function to calculate a blickhash that meets the reuired difficulty
// level of the blockchain.  it incrmenets the proof of work (POW) until
// the number of leading zeroes in the blockhash - the difficulty for the chain
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.POW++
		b.Hash = b.calculateHash()
	}
}

// Print will produce formatted output on stdout of the contents of a Block
func (b Block) Print() {
	fmt.Println("Block Info")
	fmt.Printf("Timestamp\t %s\n", b.Timestamp.String())
	fmt.Printf("Hash:\t %x\n", b.Hash)
	fmt.Printf("Prev Hash:\t %x\n", b.PreviousHash)
	fmt.Printf("POW:\t\t %d\n", b.POW)
	fmt.Printf("BlockNote:\t %s\n", b.BlockNote)
	fmt.Printf("Date:\t %s\n", b.Data) //TODO: look into the fmt of the amount field in the Map
	fmt.Println("----------")

}

// Blockchain is the list of valid blocks in the chain
type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Difficulty   int
}

// CreateBlockshain creates the genesis block for the chain and returns a instance
func CreateBlockshain(difficulty int) Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now(),
		BlockNote: "genesis block",
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// AddBlock adds a new block to the blockchain.  right now each block only holds
// one transaction
func (b *Blockchain) AddBlock(t2 transaction.Transaction) {

	blockdataJSON, _ := json.Marshal(t2)
	//fmt.Println(t2)
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := Block{
		Data:         string(blockdataJSON),
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
		BlockNote:    "heres a block note",
	}
	newBlock.mine(b.Difficulty)
	b.Chain = append(b.Chain, newBlock)

}

// IsValid walks the Blockchain to isure that none of the nodes have
// been tampered with
func (b Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		previousBlock := b.Chain[i]
		currentBlock := b.Chain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

	}
	return true
}

func Hello() {
	fmt.Println("Hello from blockchain")
}
