package main

import (
	"bytes"
)

func isChainValidCached(chain []*Block) bool {
	for i := 1; i < len(chain); i++ {
		prevBlock := chain[i-1]
		currBlock := chain[i]
		if !bytes.Equal(currBlock.PrevHash, calculateHash(prevBlock)) {
			return false
		}
		if !bytes.Equal(currBlock.Hash, calculateHash(currBlock)) {
			return false
		}
	}
	return true
}

func (bc *Blockchain) IsValid() bool {
	return isChainValidCached(bc.Blocks)
}
