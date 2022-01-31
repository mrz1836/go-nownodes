package nownodes

import "context"

// MempoolEntryResult is the mempool entry result
type MempoolEntryResult struct {
	NodeError               // Error message
	ID        string        `json:"id,omitempty"`     // Your unique ID
	Result    *MempoolEntry `json:"result,omitempty"` // Mempool entry
}

// MempoolEntry is the mempool entry
type MempoolEntry struct {
	Depends     []string `json:"depends"`
	Fee         float64  `json:"fee,omitempty"`
	Height      uint64   `json:"height"`
	ModifiedFee float64  `json:"modifiedfee,omitempty"`
	Size        int64    `json:"size,omitempty"`
	Time        int64    `json:"time"`
}

// GetMempoolEntry will get the mempool entry information for a given txID
//
// This method supports the following chains: BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC
func (c *Client) GetMempoolEntry(ctx context.Context, chain Blockchain, txID, id string) (*MempoolEntryResult, error) {

	// Validate the input
	if !chain.ValidateTxID(txID) {
		return nil, ErrInvalidTxID
	}

	// No id given?
	if len(id) == 0 {
		id = txID
	}

	// Fire the HTTP request
	results := new(MempoolEntryResult)
	if err := nodeRequest(
		ctx, c, getMempoolEntryBlockchains, chain,
		createPayload(c.options.apiKey, nodeMethodGetMempoolEntry, id, []string{txID}),
		&results,
	); err != nil {
		return nil, err
	}
	return results, nil
}
