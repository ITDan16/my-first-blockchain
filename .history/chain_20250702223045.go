package main

import (
	"time"
)

type Blockchain struct {
	Blocks     []*Block
	Difficulty int
}

func (bc *Blockchain) AddBlock(data string) *Block {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:     lastBlock.Index + 1,
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  lastBlock.Hash,
	}
	hash, nonce := proofOfWork(newBlock, bc.Difficulty)
	newBlock.Hash = hash
	newBlock.Nonce = nonce
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}
