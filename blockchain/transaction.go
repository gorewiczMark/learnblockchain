package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int    //Amount of coins in transaction
	PubKey string //Public key to identify who sent the coins
}

type TxInput struct {
	ID  []byte //ID of the transaction that an output is inside
	Out int    //Index of the specific output we found within a transaction
	Sig string //Should be script to add data to output's pubkey but for now is same as pubkey
}

func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}

	txIn := TxInput{[]byte{}, -1, data}

	txOut := TxOutput{reward, toAddress}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (tx *Transaction) IsCoinBase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	//STEP 1
	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	//STEP 2
	if acc < amount {
		log.Panic("Error: Not enough funds!")
	}

	//STEP 3
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	//STEP 4
	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	//STEP 5
	tx := Transaction{nil, inputs, outputs}
	//STEP 6
	tx.SetID()

	return &tx
}
