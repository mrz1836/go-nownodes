package nownodes

import (
	"context"
)

// TransactionInfo is the transaction information returned to the GetTransaction request
type TransactionInfo struct {
	BlockHash     string    `json:"blockHash"`
	BlockHeight   int64     `json:"blockHeight"`
	BlockTime     int64     `json:"blockTime"`
	Confirmations int64     `json:"confirmations"`
	Fees          string    `json:"fees"`
	Hex           string    `json:"hex"`
	LockTime      int64     `json:"lockTime,omitempty"` // BTC
	TxID          string    `json:"txid"`
	Value         string    `json:"value"`
	ValueIn       string    `json:"valueIn"`
	Version       int8      `json:"version"`
	Vin           []*Input  `json:"vin"`
	VOut          []*Output `json:"vout"`
}

// Input is the transaction input
type Input struct {
	Addresses []string `json:"addresses,omitempty"`
	Coinbase  string   `json:"coinbase,omitempty"`
	Hex       string   `json:"hex,omitempty"`
	IsAddress bool     `json:"isAddress"`
	N         int64    `json:"n"`
	Sequence  int64    `json:"sequence"`
	TxID      string   `json:"txid,omitempty"`
	Value     string   `json:"value,omitempty"`
	VOut      int64    `json:"vout"`
}

// Output is the transaction output
type Output struct {
	Addresses []string `json:"addresses,omitempty"`
	Hex       string   `json:"hex,omitempty"`
	IsAddress bool     `json:"isAddress"`
	N         int64    `json:"n"`
	Spent     bool     `json:"spent"`
	Value     string   `json:"value,omitempty"`
}

// GetTransaction will get transaction information by a given TxID
//
// This method supports the following chains: BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC
func (c *Client) GetTransaction(ctx context.Context, chain Blockchain, txID string) (*TransactionInfo, error) {

	// Validate the input
	if !chain.ValidateTx(txID) {
		return nil, ErrInvalidTxID
	}

	// Fire the HTTP request
	info := new(TransactionInfo)
	if err := fireBlockBookRequest(
		ctx, c, getTransactionBlockchains, chain, routeGetTx+txID, &info,
	); err != nil {
		return nil, err
	}
	return info, nil
}
