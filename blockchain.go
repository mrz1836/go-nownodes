package nownodes

import (
	"encoding/hex"
	"strings"
)

// Blockchain is the supported blockchain networks
type Blockchain string

// Supported blockchains
const (
	BCH        Blockchain = blockchainBCH        // BitcoinCash: https://bch.info/
	BSV        Blockchain = blockchainBSV        // BitCoin: https://bitcoinsv.com
	BTC        Blockchain = blockchainBTC        // BitCore: https://bitcoin.org
	BTCTestnet Blockchain = blockchainBTCTestnet // BitCore Testnet: https://bitcoin.org
	BTG        Blockchain = blockchainBTG        // BitGold: https://bitcoingold.org/
	DASH       Blockchain = blockchainDASH       // Dash: https://www.dash.org/
	DOGE       Blockchain = blockchainDOGE       // DogeCoin: https://dogecoin.com/
	LTC        Blockchain = blockchainLTC        // LiteCoin: https://litecoin.org/
)

// String is the string version of the blockchain
func (n Blockchain) String() string {
	return string(n)
}

// BlockBookURL is the url for the block book API
func (n Blockchain) BlockBookURL() string {
	switch n {
	case BCH:
		return blockchainBCH + "." + nowNodesURL
	case BSV:
		return blockchainBSV + "." + nowNodesURL
	case BTC:
		return blockchainBTC + "." + nowNodesURL
	case BTCTestnet:
		return blockchainBTCTestnet + "." + nowNodesURL
	case BTG:
		return blockchainBTG + "." + nowNodesURL
	case DASH:
		return blockchainDASH + "." + nowNodesURL
	case DOGE:
		return blockchainDOGE + "." + nowNodesURL
	case LTC:
		return blockchainLTC + "." + nowNodesURL
	default:
		return ""
	}
}

// ValidateTxID will do basic validations on the tx id string
func (n Blockchain) ValidateTxID(txID string) bool {
	switch n {
	case BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC:
		return len(txID) == bitcoinTransactionLength
	default:
		return false
	}
}

// ValidateTxHex will do basic validations on the tx hex string
func (n Blockchain) ValidateTxHex(txHex string) bool {
	switch n {
	case BCH, BSV, BTC, BTCTestnet, BTG, DASH, DOGE, LTC:
		if b, err := hex.DecodeString(
			txHex,
		); err != nil || len(b) == 0 {
			return false
		}
		return true
	default:
		return false
	}
}

// ValidateAddress will do basic validations on the address
func (n Blockchain) ValidateAddress(address string) bool {
	switch n {
	case BCH:
		// note: validate that it's a LTC address (prefix)
		withoutPrefix := strings.ReplaceAll(address, bitcoinCashPrefix, "")
		return len(withoutPrefix) >= bitcoinMinAddressLength && len(withoutPrefix) <= bitcoinCashMaxAddressLength
	case BSV:
		return len(address) >= bitcoinMinAddressLength && len(address) <= bitcoinMaxAddressLength
	case BTC, BTCTestnet, BTG, DOGE:
		return len(address) >= bitcoinMinAddressLength && len(address) <= bitcoinMaxAddressLength
	case DASH:
		// note: validate that it's a DASH address (prefix)
		return len(address) >= bitcoinMinAddressLength && len(address) <= bitcoinMaxAddressLength
	case LTC:
		// note: validate that it's a LTC address (prefix)
		return len(address) >= bitcoinMinAddressLength && len(address) <= liteCoinMaxAddressLength
	default:
		return false
	}
}

// isBlockchainSupported will return true if the blockchain was found in the list
func isBlockchainSupported(list []Blockchain, blockchain Blockchain) bool {
	for _, chain := range list {
		if chain == blockchain {
			return true
		}
	}
	return false
}
