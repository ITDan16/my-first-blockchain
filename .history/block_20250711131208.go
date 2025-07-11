package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
)

type Block struct {
	Index                 int
	Timestamp             int64
	Data                  []byte
	PrevHash              []byte
	Hash                  []byte
	Nonce                 int
	explicitlyInitialized bool
}

func serializeBlock(block *Block) []byte {
	var buf bytes.Buffer

	buf.WriteByte(0x01)
	if block.explicitlyInitialized {
		buf.WriteByte(0xFF)
	} else {
		buf.WriteByte(0x00)
	}

	_ = binary.Write(&buf, binary.LittleEndian, int64(block.Index))
	_ = binary.Write(&buf, binary.LittleEndian, int64(block.Timestamp))
	_ = binary.Write(&buf, binary.LittleEndian, int64(block.Nonce))

	_ = binary.Write(&buf, binary.LittleEndian, int32(len(block.Data)))
	buf.Write(block.Data)

	_ = binary.Write(&buf, binary.LittleEndian, int32(len(block.PrevHash)))
	buf.Write(block.PrevHash)

	return buf.Bytes()
}

func calculateHash(block *Block) []byte {
	bytes := serializeBlock(block)
	hash := sha256.Sum256(bytes)
	return hash[:]
}

func proofOfWork(block *Block, difficulty int) ([]byte, int) {
	prefix := strings.Repeat("0", difficulty)
	nonce := 0
	var hash []byte
	for {
		block.Nonce = nonce
		hash = calculateHash(block)
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			break
		}
		nonce++
	}
	return hash, nonce
}
