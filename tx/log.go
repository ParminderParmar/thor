package tx

import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/vechain/thor/thor"
	"github.com/vechain/thor/trie"
)

// Log represents a contract log event. These events are generated by the LOG opcode and
// stored/indexed by the node.
type Log struct {
	// address of the contract that generated the event
	Address thor.Address
	// list of topics provided by the contract.
	Topics []thor.Bytes32
	// supplied by the contract, usually ABI-encoded
	Data []byte
}

// Logs slice of logs.
type Logs []*Log

// RootHash computes merkle root hash of receipts.
func (ls Logs) RootHash() thor.Bytes32 {
	if len(ls) == 0 {
		// optimized
		return emptyRoot
	}
	return trie.DeriveRoot(derivableLogs(ls))
}

// implements DerivableList
type derivableLogs Logs

func (ls derivableLogs) Len() int {
	return len(ls)
}
func (ls derivableLogs) GetRlp(i int) []byte {
	data, err := rlp.EncodeToBytes(ls[i])
	if err != nil {
		panic(err)
	}
	return data
}
