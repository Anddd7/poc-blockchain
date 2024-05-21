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

func (c *Chain) VerifyBlock(block Block) bool {
	// check block is valid, e.g. hash, transactions
	// check parent is existing
	parent := c.Find(block.PreHash)
	return parent != nil
}

// Find find a block by hash
func (c *Chain) Find(hash Hash) *Block {
	for _, block := range c.Blocks {
		if block.Hash == hash {
			return &block
		}
	}
	return nil
}

// AddBlock add a new block to the chain after broadcast is successful
func (c *Chain) AddBlock(block Block) {
	// if we got multiple blocks at the same time, somehow conflict will occur
	
	// * keep longest chain, choose the block that has nearest parent
	// A -- B -- C
	// 	\-- D
	// while C and D are broadcasted at the same time, we choose C

	// * invalid block will be dropped
	// but, if that block is already in local chain, we'll replace/repair the chain
	// A -- B -- C
	//  \-- D -- E -- F
	// D,E,F has been approved in the network, so the B -- C needs to be rollbacked


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
