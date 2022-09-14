package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hase"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hashbytes := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hashbytes)
}

func getLastHash() string {
	totalLength := len(GetBlockChain().blocks)
	if totalLength == 0 {
		return ""
	}
	lastBlock := b.blocks[totalLength-1]
	return lastBlock.Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

var ErrorNotFound = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrorNotFound
	}
	return b.blocks[height-1], nil
}
