package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Block the data structure that represents the block in the blockchain
type Block struct {
	Hash         Hash
	PreHash      Hash
	Transactions []Transaction
	Metadata     Metadata
	Timestamp    time.Time
}

// Hash the indentifier of the block
type Hash struct {
	Value string
}

// Metadata the extensible data of the block
type Metadata map[string]string

// Transaction the transaction data structure
type Transaction struct {
	From  Wallet
	To    Wallet
	Value Money
}

// Money the money data structure
type Money struct {
	Amount int
}

// Wallet the wallet data structure
type Wallet struct {
	Address string
}

// toString capture and summarize the block data
func (b *Block) toString() (string, error) {
	bytes, err := json.Marshal(b)
	return string(bytes), err
}

// InitGenesisBlock create a genesis block
func InitGenesisBlock() Block {
	return Block{
		Hash:      Hash{Value: fmt.Sprintf("%064x", 0)},
		Metadata:  Metadata{"name": "Genesis Block"},
		Timestamp: time.Now(),
	}
}
