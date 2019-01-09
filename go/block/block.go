package block

import (
	"crypto/sha256"
	"strconv"
	"time"
)

// Block is a atomic unit of a blockchain
type Block struct {
	Index     int
	Timestamp int64
	Hash      string
	PrevHash  string
	Data      string
}

// CalculateHash returns the string hash generated by SHA256
func CalculateHash(index int, timestamp int64, data, prevHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(strconv.Itoa(index)))
	hasher.Write([]byte(string(timestamp)))
	hasher.Write([]byte(data))
	hasher.Write([]byte(prevHash))
	return string(hasher.Sum(nil))
}

// New creates instance of Block
func New(index int, data, prevHash string) *Block {
	timestamp := time.Now().Unix()
	hash := CalculateHash(index, timestamp, data, prevHash)
	return &Block{
		index,
		timestamp,
		hash,
		prevHash,
		data,
	}
}
