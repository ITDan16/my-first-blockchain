package main

import (
    "bytes"
    "errors"
    "fmt"
)

var (
    ErrInvalidPrevHash = errors.New("previous hash mismatch")
    ErrInvalidHash     = errors.New("invalid block hash")
    ErrInvalidPOW      = errors.New("block does not satisfy proof-of-work requirement")
)

// Валидация блока с деталями (и с учётом сложности PoW)
func ValidateBlock(prev, curr *Block, config Config) error {
    if !bytes.Equal(curr.PrevHash, prev.Hash) {
        return ErrInvalidPrevHash
    }
    if !bytes.Equal(curr.Hash, calculateHash(curr)) {
        return ErrInvalidHash
    }
    // Проверка сложности PoW
    prefix := ""
    if config.Difficulty > 0 {
        prefix = fmt.Sprintf("%0*s", config.Difficulty, "")
    }
    if config.Difficulty > 0 && !bytes.HasPrefix([]byte(fmt.Sprintf("%x", curr.Hash)), []byte(prefix)) {
        return ErrInvalidPOW
    }
    return nil
}