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
	N         uint64   `json:"n"`
	Sequence  int64    `json:"sequence"`
	TxID      string   `json:"txid,omitempty"`
	Value     string   `json:"value,omitempty"`
	VOut      uint64   `json:"vout"`
}

// Output is the transaction output
type Output struct {
	Addresses []string `json:"addresses,omitempty"`
	Hex       string   `json:"hex,omitempty"`
	IsAddress bool     `json:"isAddress"`
	N         uint64   `json:"n"`
	Spent     bool     `json:"spent"`
	Value     string   `json:"value,omitempty"`
}

// BroadcastResult is the successful broadcast results
type BroadcastResult struct {
	Result string `json:"result"` // {"result": "15e78db3a6247ca320de2202240f6a4877ea3af338e23bf5ff3e5cbff3763bf6"}
}

// GetTransaction will get transaction information by a given TxID
//
// This method supports the following chains: BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC
func (c *Client) GetTransaction(ctx context.Context, chain Blockchain, txID string) (*TransactionInfo, error) {

	// Validate the input
	if !chain.ValidateTxID(txID) {
		return nil, ErrInvalidTxID
	}

	// Fire the HTTP request
	info := new(TransactionInfo)
	if err := blockBookRequest(
		ctx, c, getTransactionBlockchains, chain, routeGetTx+txID, &info,
	); err != nil {
		return nil, err
	}
	return info, nil
}

// SendTransaction will submit a broadcast request with the given tx hex payload
//
// This method supports the following chains: BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC
func (c *Client) SendTransaction(ctx context.Context, chain Blockchain, txHex string) (*BroadcastResult, error) {

	// Validate the input
	if !chain.ValidateTxHex(txHex) {
		return nil, ErrInvalidTxHex
	}

	// Max size of a GET request: 2048 (not sure how NowNodes is handling this)
	if len(txHex) > maxTxHexLengthOnSend {
		return nil, ErrTxHexTooLarge
	}

	// Fire the HTTP request
	result := new(BroadcastResult)
	if err := blockBookRequest(
		ctx, c, sendTransactionBlockchains, chain, routeSendTx+txHex, &result,
	); err != nil {
		return nil, err
	}
	return result, nil
}
