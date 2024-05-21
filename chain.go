package main

// Chain the chain of blocks
type Chain struct {
	Difficulty int
	Blocks     []Block
}

// NewChain create a new chain with genesis block
func NewChain(difficulty int) Chain {
	return Chain{
		Difficulty: difficulty,
		Blocks:     []Block{InitGenesisBlock()},
	}
}

// AddBlock add a new block to the chain after broadcast is successful
func (c *Chain) AddBlock(block Block) {
	c.Blocks = append(c.Blocks, block)
}

// LatestHash get the latest hash of the chain
func (c *Chain) LatestHash() Hash {
	return c.Blocks[len(c.Blocks)-1].Hash
}

// Balance get the balance of the wallet
// UTXO: calculate the balance by iterating through all transactions
func (c *Chain) Balance(wallet Wallet) Money {
	balance := 0
	for _, block := range c.Blocks {
		for _, tx := range block.Transactions {
			if tx.From.Address == wallet.Address {
				balance = balance - tx.Value.Amount
			}
			if tx.To.Address == wallet.Address {
				balance = balance + tx.Value.Amount
			}
		}
	}

	return Money{Amount: balance}
}
