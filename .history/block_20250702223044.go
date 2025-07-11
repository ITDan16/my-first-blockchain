package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func serializeBlock(block *Block) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, int64(block.Index))
	binary.Write(&buf, binary.LittleEndian, int64(block.Timestamp))
	binary.Write(&buf, binary.LittleEndian, int64(block.Nonce))
	binary.Write(&buf, binary.LittleEndian, int32(len(block.Data)))
	buf.Write(block.Data)
	binary.Write(&buf, binary.LittleEndian, int32(len(block.PrevHash)))
	buf.Write(block.PrevHash)
	return buf.Bytes()
}

func calculateHash(block *Block) []byte {
	bytes := serializeBlock(block)
	hash := sha256.Sum256(bytes)
	return hash[:]
}
