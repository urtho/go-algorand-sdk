package models

import "github.com/urtho/go-algorand-sdk/v2/types"

// SimulateRequestTransactionGroup a transaction group to simulate.
type SimulateRequestTransactionGroup struct {
	// Txns an atomic transaction group.
	Txns []types.SignedTxn `json:"txns"`
}
