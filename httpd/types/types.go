package types

// UTXO models the data corresponding to unspent transaction outputs.
// Convenience type; for limited use only.
type UTXO struct {
	Value   int64
	Address string
}

// Input models data corresponding to transaction inputs.
type Input struct {
	Coinbase    string   `json:"coinbase,omitempty"`         // [coinbase] The coinbase encoded as hex
	OutputHash  string   `json:"output_hash,omitempty"`      // [non-coinbase] Same as transaction ID of vin
	OutputIndex uint32   `json:"output_index,omitempty"`     // [non-coinbase] Index of the corresponding UTXO
	Value       int64    `json:"value,omitempty"`            // [non-coinbase] Value of the corresponding UTXO in satoshis
	Address     string   `json:"address,omitempty"`          // [non-coinbase] Address of the corresponding UTXO; can be empty
	ScriptSig   string   `json:"script_signature,omitempty"` // [non-coinbase] Hex-encoded signature script
	Witness     []string `json:"txinwitness,omitempty"`      // [non-coinbase] Array of hex-encoded witness data
	InputIndex  int      `json:"input_index"`                // [all] Non-standard data required by Ledger Blockchain Explorer
	Sequence    uint32   `json:"sequence"`                   // [all] Input sequence number, used to track unconfirmed txns
}

// Output models data corresponding to transaction outputs.
type Output struct {
	OutputIndex uint32 `json:"output_index"`      // Used to uniquely identify an output in a transaction
	Value       int64  `json:"value"`             // Value of output in satoshis
	ScriptHex   string `json:"script_hex"`        // Hex-encoded script
	Address     string `json:"address,omitempty"` // Address of the UTXO; can be empty
}

// SparseBlock models data corresponding to a block, but with limited information.
// It is used to represent minimal information of the block containing the given
// transaction.
type SparseBlock struct {
	Hash   string `json:"hash"`
	Height int64  `json:"height"`
	Time   string `json:"time"`
}

// Transaction represents the principal type to model the response of the GetTransaction handler.
type Transaction struct {
	ID            string      `json:"id"`
	Hash          string      `json:"hash"`
	ReceivedAt    string      `json:"received_at"`
	LockTime      uint32      `json:"lock_time"`
	Fees          int64       `json:"fees"`
	Confirmations uint64      `json:"confirmations"`
	Inputs        []Input     `json:"inputs"`
	Outputs       []Output    `json:"outputs"`
	Block         SparseBlock `json:"block"`
}