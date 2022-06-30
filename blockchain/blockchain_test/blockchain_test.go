package blockchain_test

import (
	"reflect"
	"testing"

	. "github.com/gorewiczMark/learnblockchain/blockchain"
)

func TestInitBlockChain(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want *BlockChain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitBlockChain(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitBlockChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
