package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Used before Proof of Work Implemented to create blocks
// func (b *Block) DeriveHash() {
// 	// This will join our previous block's relevant info with the new block
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

// 	//This preforms the actual hashing algoritm
// 	hash := sha256.Sum256(info)

// 	b.Hash = hash[:]
// }

// CreateBlock
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// TODO: look at improving error handling
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}
