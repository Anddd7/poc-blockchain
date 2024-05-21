package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Node the agent that hosts blockchain network
type Node struct {
	Chain     Chain
	Mempool   []Transaction
	Connected []Node
}

// SubmitTransaction create a new transaction in current node
func (n *Node) SubmitTransaction(from Wallet, to Wallet, value Money) error {
	n.Mempool = append(n.Mempool, Transaction{
		From:  from,
		To:    to,
		Value: value,
	})

	return nil
}

// PackageBlock package current transactions into a block
func (n *Node) PackageBlock(name string, metadata Metadata) Block {
	metadata["name"] = name

	return Block{
		PreHash:      n.Chain.LatestHash(),
		Transactions: n.Mempool,
		Metadata:     metadata,
		Timestamp:    time.Now(),
	}
}

// Nonce "mine" the block by finding the nonce
func (n *Node) Nonce(block Block) (Block, error) {
	value, err := block.toString()
	if err != nil {
		return Block{}, err
	}

	// simulate the proof of work, high difficulty means longer time to compute
	// when multiple nodes are competing to mine the block
	// only the node with the highest computation power will win
	hash := ""
	nonce := 0
	for !strings.HasPrefix(hash, strings.Repeat("0", n.Chain.Difficulty)) {
		nonce = nonce + 1
		hash = n.hash(value + fmt.Sprintf("%d", nonce))
	}

	block.Hash.Value = hash

	return block, nil
}

// hash the proof of work function
func (n *Node) hash(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Broadcast the block to the network
func (n *Node) Broadcast(block Block) error {
	if n.Chain.VerifyBlock(block) {
		n.Chain.AddBlock(block)
		n.Mempool = []Transaction{}
	}

	// broadcast to other node
	approved := 0
	for _, node := range n.Connected {
		if node.Chain.VerifyBlock(block) {
			approved++
		}
	}

	// if less than half of the network approved the block, rollback
	if approved < len(n.Connected)/2 {
		// rollback the block
		n.Chain.Blocks = n.Chain.Blocks[:len(n.Chain.Blocks)-1]
		n.Mempool = append(n.Mempool, block.Transactions...)
	}

	return nil
}

// ReceiveBlock receive a block from the network
func (n *Node) ReceiveBlock(block Block) error {
	// the node will try to accept the block and add to the chain
	// but sometimes, the validation will fail - forked chain, invalid block, etc
	if n.Chain.VerifyBlock(block) {
		n.Chain.AddBlock(block)
	}

	// if you got a forked chain, you need to choose the longest chain
	// - find the nearest and in-chain parent
	// -- if the parent is not in local chain, you need to request more info from the network
	// -- check the height, pick the longest chain
	// - drop and rollback the invalid chain
	// -- re-queue the transactions
	// -- re-package the block

	return nil
}
