package nownodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain_String(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		chain    Blockchain
		expected string
	}{
		{BCH, blockchainBCH},
		{BSV, blockchainBSV},
		{BTC, blockchainBTC},
		{BTCTestnet, blockchainBTCTestnet},
		{BTG, blockchainBTG},
		{DASH, blockchainDASH},
		{DOGE, blockchainDOGE},
		{LTC, blockchainLTC},
	}

	for _, testCase := range tests {
		t.Run("chain "+testCase.chain.String()+": String()", func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.chain.String())
		})
	}

	t.Run("unknown blockchain", func(t *testing.T) {
		u := Blockchain("unknown")
		assert.Equal(t, "unknown", u.String())
	})
}

func TestBlockchain_BlockBookURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		chain    Blockchain
		expected string
	}{
		{BCH, blockchainBCH + "." + nowNodesURL},
		{BSV, blockchainBSV + "." + nowNodesURL},
		{BTC, blockchainBTC + "." + nowNodesURL},
		{BTCTestnet, blockchainBTCTestnet + "." + nowNodesURL},
		{BTG, blockchainBTG + "." + nowNodesURL},
		{DASH, blockchainDASH + "." + nowNodesURL},
		{DOGE, blockchainDOGE + "." + nowNodesURL},
		{LTC, blockchainLTC + "." + nowNodesURL},
	}

	for _, testCase := range tests {
		t.Run("chain "+testCase.chain.String()+": BlockBookURL()", func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.chain.BlockBookURL())
		})
	}

	t.Run("unknown blockchain", func(t *testing.T) {
		u := Blockchain("unknown")
		assert.Empty(t, u.BlockBookURL())
	})
}

func TestBlockchain_ValidateTx(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		chain    Blockchain
		txID     string
		expected bool
	}{
		{BCH, testTxID(BCH), true},
		{BSV, testTxID(BSV), true},
		{BTC, testTxID(BTC), true},
		{BTCTestnet, testBTCTxID, true},
		{BTG, testTxID(BTG), true},
		{DASH, testTxID(DASH), true},
		{DOGE, testTxID(DOGE), true},
		{LTC, testTxID(LTC), true},
		{BSV, "", false},
		{BSV, "12345", false},
		{BSV, testBitcoinTxID + "1", false},
	}

	for _, testCase := range tests {
		t.Run("chain "+testCase.chain.String()+": ValidateTx()", func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.chain.ValidateTx(testCase.txID))
		})
	}

	t.Run("unknown blockchain", func(t *testing.T) {
		u := Blockchain("unknown")
		assert.Equal(t, false, u.ValidateTx(testBitcoinTxID))
		assert.Equal(t, false, u.ValidateTx(""))
	})
}
