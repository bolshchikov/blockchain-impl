package blockchain

import (
	"blockchain-impl/go/block"
)

// Blockchain is a group of blocks linked together
type Blockchain struct {
	difficulty int
	chain      []*block.Block
	length     int
}

// Len returns the length of the blockchain
func (bc *Blockchain) Len() int {
	return bc.length
}

// AddBlock creates a block with given data and adds it to the chain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := *bc.GetLastBlock()
	newBlock := block.New(
		bc.length,
		data,
		prevBlock.Hash,
	)
	newBlock.MineBlock(bc.difficulty)
	bc.chain = append(bc.chain, newBlock)
	bc.length++
}

// GetLastBlock returns the pointer to the last block in the blockchain
func (bc *Blockchain) GetLastBlock() *block.Block {
	return bc.chain[bc.length-1]
}

func addGenesisBlock(bc *Blockchain) {
	genesis := block.New(0, "", "")
	bc.chain = append(bc.chain, genesis)
	bc.length++
}

// New returns an instance of a blockchain
func New(difficulty int) *Blockchain {
	bc := Blockchain{
		difficulty: difficulty,
	}
	addGenesisBlock(&bc)
	return &bc
}

// IsValidChain checks the validity of a chain based on hashes
func IsValidChain(bc *Blockchain) bool {
	for i := 1; i < bc.length; i++ {
		currBlock := *bc.chain[i]
		prevBlock := *bc.chain[i-1]
		currHash := currBlock.CalculateHash()
		if currBlock.Hash != currHash {
			return false
		}
		if prevBlock.Hash != currBlock.PrevHash {
			return false
		}
		return true
	}
	return true
}
