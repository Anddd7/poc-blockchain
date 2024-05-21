package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, blockchain")

	node := Node{
		Chain:      Chain{},
		LatestHash: Hash{GENESIS_NONCE},
	}

	printJson("created node", node)

	transactions := []Transaction{
		{
			From:  Wallet{Address: "me", Balance: Money{Amount: 100}},
			To:    Wallet{Address: "you", Balance: Money{Amount: 0}},
			Value: Money{Amount: 100},
		},
	}
	block := node.NewBlock(transactions, Metadata{
		"timestamp": time.Now().String(),
	})

	printJson("created block", block)

	block, _ = node.Nonce(block)

	printJson("mined block", block)

	_ = node.Broadcast(block)

	printJson("broadcasted block", block)
}

func printJson(msg string, obj interface{}) {
	json, _ := json.Marshal(obj)
	fmt.Printf("%s: %v\n", msg, string(json))
}
