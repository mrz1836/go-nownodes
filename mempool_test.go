package nownodes

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetMempoolEntry(t *testing.T) {
	t.Parallel()

	t.Run("valid cases", func(t *testing.T) {

		var tests = []struct {
			chain      Blockchain
			txID       string
			id         string
			expectedID string
		}{
			{BCH, testTxID(BCH), testUniqueID, testUniqueID},
			{BSV, testTxID(BSV), testUniqueID, testUniqueID},
			{BSV, testTxID(BSV), "", testTxID(BSV)},
			{BTC, testTxID(BTC), testUniqueID, testUniqueID},
			{BTCTestnet, testTxID(BTCTestnet), testUniqueID, testUniqueID},
			{BTG, testTxID(BTG), testUniqueID, testUniqueID},
			{DASH, testTxID(DASH), testUniqueID, testUniqueID},
			{DOGE, testTxID(DOGE), testUniqueID, testUniqueID},
			{LTC, testTxID(LTC), testUniqueID, testUniqueID},
		}

		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validNodeResponse{}))
		ctx := context.Background()

		for _, testCase := range tests {
			t.Run("chain "+testCase.chain.String()+": GetMempoolEntry("+testCase.txID+")", func(t *testing.T) {
				results, err := c.GetMempoolEntry(ctx, testCase.chain, testCase.txID, testCase.id)
				require.NoError(t, err)
				require.NotNil(t, results)

				assert.Equal(t, testCase.expectedID, results.ID)
				assert.Greater(t, results.Result.Time, int64(0))
				assert.Greater(t, results.Result.Fee, float64(0))
			})
		}
	})

	t.Run("missing or invalid tx id", func(t *testing.T) {
		type testData struct {
			chain Blockchain
			txID  string
			id    string
			err   error
		}
		var testCases []testData
		for _, chain := range sendTransactionBlockchains {
			testCases = append(testCases, testData{chain: chain, txID: "", id: testUniqueID, err: ErrInvalidTxID})
			testCases = append(testCases, testData{chain: chain, txID: "12345", id: testUniqueID, err: ErrInvalidTxID})
			testCases = append(testCases, testData{chain: chain, txID: "invalid-tx-hex", id: testUniqueID, err: ErrInvalidTxID})
		}

		c := NewClient(WithHTTPClient(&validNodeResponse{}))
		ctx := context.Background()

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": invalid-tx", func(t *testing.T) {
				results, err := c.GetMempoolEntry(ctx, testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
				assert.ErrorIs(t, err, testCase.err)
			})
		}
	})

	t.Run("unsupported chain", func(t *testing.T) {
		c := NewClient(WithHTTPClient(&validNodeResponse{}))
		ctx := context.Background()
		getMempoolEntryBlockchains = []Blockchain{BSV}
		results, err := c.GetMempoolEntry(ctx, BTC, testTxID(BTC), testUniqueID)
		require.Error(t, err)
		require.Nil(t, results)
		assert.ErrorIs(t, err, ErrUnsupportedBlockchain)
	})

	t.Run("error cases", func(t *testing.T) {

		type testData struct {
			chain Blockchain
			txID  string
			id    string
		}
		var testCases []testData
		for _, chain := range getMempoolEntryBlockchains {
			testCases = append(testCases, testData{chain: chain, txID: testTxHex(chain), id: testUniqueID})
		}

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": broadcast error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorNodeErrorResponse{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": http req error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqErr{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": missing body contents", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqNoBodyErr{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": error with resp", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqWithRespErr{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": invalid json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadJSONResponse{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": invalid error json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadErrorJSONResponse{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": missing api key", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorMissingAPIKey{}))
				results, err := c.GetMempoolEntry(context.Background(), testCase.chain, testCase.txID, testCase.id)
				require.Error(t, err)
				require.Nil(t, results)
			})
		}
	})
}

func ExampleClient_GetMempoolEntry() {
	c := NewClient(WithHTTPClient(&validNodeResponse{}))
	info, _ := c.GetMempoolEntry(context.Background(), BSV, testTxID(BSV), testUniqueID)
	fmt.Printf("tx in mempool time: %d", info.Result.Time)
	// Output:tx in mempool time: 1643661192
}

func BenchmarkClient_GetMempoolEntry(b *testing.B) {
	c := NewClient(WithHTTPClient(&validNodeResponse{}))
	ctx := context.Background()
	tx := testTxID(BSV)
	for i := 0; i < b.N; i++ {
		_, _ = c.GetMempoolEntry(ctx, BSV, tx, testUniqueID)
	}
}
