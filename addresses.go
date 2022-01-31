package nownodes

import (
	"context"
)

// AddressInfo is the address information returned to the GetAddress request
type AddressInfo struct {
	Address            string   `json:"address"`
	Balance            string   `json:"balance"`
	ItemsOnPage        uint64   `json:"itemsOnPage"`
	Page               uint64   `json:"page"`
	TotalPages         uint64   `json:"totalPages"`
	TotalReceived      string   `json:"totalReceived"`
	TotalSent          string   `json:"totalSent"`
	TxIDs              []string `json:"txids,omitempty"`
	Txs                uint64   `json:"txs"`
	UnconfirmedBalance string   `json:"unconfirmedBalance"`
	UnconfirmedTxs     uint64   `json:"unconfirmedTxs"`
}

// GetAddress will get address information by a given address
//
// This method supports the following chains: BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC
func (c *Client) GetAddress(ctx context.Context, chain Blockchain, address string) (*AddressInfo, error) {

	// Validate the input
	if !chain.ValidateAddress(address) {
		return nil, ErrInvalidAddress
	}

	// Fire the HTTP request
	info := new(AddressInfo)
	if err := blockBookRequest(
		ctx, c, getAddressBlockchains, chain, routeGetAddress+address, &info,
	); err != nil {
		return nil, err
	}
	return info, nil
}
