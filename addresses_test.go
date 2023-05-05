package nownodes

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// validAddressResponse will return a valid address for all supported blockchains
type validAddressResponse struct{}

func (v *validAddressResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	// Address data
	addressResponses := map[string]string{
		BCH.String():  `{"page":1,"totalPages":104,"itemsOnPage":1000,"address":"bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn","balance":"3706237","totalReceived":"1053381020454","totalSent":"1053377314217","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":103219,"txids":["fc2ee06ff4a22a630db1222ac1afa3595046854d03b12e6f423ee37e7db52be8","640f5f203691199838efb0ab43a5c9f15cadbf4c573523058816f02590f95dff","ac8e25de62f92dd2143c85b5487d2b77a34954c564bfe3047bc34639711c8b63"]}`,
		BSV.String():  `{"page":1,"totalPages":174,"itemsOnPage":1000,"address":"1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL","balance":"101556","totalReceived":"66351012","totalSent":"66249456","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":173661,"txids":["12d9e1ed43444f03dee7381f66ecc206fe14719523a72d2c6536e1f4c4e05a14","bc762632e1a02417ba0ee69b7d19059d3ba4f7905320f3930f22be6e9dde8610","fcf4cda23a9394d446a421a15bfd7da0095fd701f2fc3c446d6a54e744caf1e2","baee5d90e542b146eb3ce65e2b9d5f39c3a178b4134c583a6c9e379162c7604c","b1e1d62e2ad74aab16741779bae5c03b64d35dba176687f88e475b9b5504fd6e"]}`,
		BTC.String():  `{"page":1,"totalPages":1,"itemsOnPage":1000,"address":"17eKje3fzPs633GKsotMFkLKRzv1HPSRTz","balance":"4756408","totalReceived":"209893262","totalSent":"205136854","unconfirmedBalance":"-854392","unconfirmedTxs":9,"txs":217,"txids":["bf8892afca454f5822dfe51eb8aebabab64b2803ea5ff8902b1bc6b0b1dd5e08","e9da28382a504a1f86b063cc5b4c4d242af006ec61f4189310880da8f387ab3f","3f696031a6e2a07364c5d5e6d829f3fd85fcf6f91f8b6fa7a4d0f24b563f3373"]}`,
		BTG.String():  `{"page":1,"totalPages":13,"itemsOnPage":1000,"address":"ATTav2PtmotBZwxgWjrZCgaZpE89kcJ29B","balance":"121734541552","totalReceived":"14394650955688","totalSent":"14272916414136","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":12287,"txids":["f52056973f853356d059e2cd05e214e8c405cab398240e811a54b622eff61c56","ca5a3bbd6245598ec1af177b2d56e370139d96217600dec1455216ae87211314","0ba336fe89ca26d04aaeabed8b2b252dba45004c389216af15290b6ff9179737"]}`,
		DASH.String(): `{"page":1,"totalPages":1,"itemsOnPage":1000,"address":"Xe7tPVUvDpt52h2KykMpj4hmh8VsAaPCgt","balance":"0","totalReceived":"54100000","totalSent":"54100000","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":2,"txids":["7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e","b9e4cebbf1cde3118ec62df72553a7d60fad78565028e8bf1c31edec23a1a393"]}`,
		DOGE.String(): `{"page":1,"totalPages":18,"itemsOnPage":1000,"address":"ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a","balance":"16765841411433","totalReceived":"189519869773689377","totalSent":"189503103932277944","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":17885,"txids":["9a3bf978b4e48251f7c5c640afc6cf5cf9f0888b7d9cfca3b4db65d14b99bccf","7b36a2713b5b3e12633a2a989dff177d31a0d49df50010e8822a5ce0ee70ab3d","c65acb715eec3eec0dd405ff0da566ee6b5a7af3c2cfc87a2c4d87a7e1c1f01e"]}`,
		LTC.String():  `{"page":1,"totalPages":1,"itemsOnPage":1000,"address":"ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps","balance":"4183327","totalReceived":"97526951","totalSent":"93343624","unconfirmedBalance":"0","unconfirmedTxs":0,"txs":32,"txids":["72c9a6bc19049c6f455477af8561a726352a61f55c5a16f8165f09503951d94d","1de19db94cab8f7bc4aaf15691c03732ed656ae8ce374e9b3d347520b563c49b","2381fc53f03717267895a885109075f8ba275401fbea13f9e1ab8994c3017889"]}`,
	}

	// Valid response
	for _, chain := range getAddressBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeGetAddress+testAddress(chain)) {
			resp.StatusCode = http.StatusOK
			resp.Body = io.NopCloser(bytes.NewBuffer([]byte(addressResponses[chain.String()])))
			return resp, nil
		}
	}

	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

// errorInvalidAddress will return an error for the address response
type errorInvalidAddress struct{}

func (v *errorInvalidAddress) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	// Error response
	for _, chain := range getAddressBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeGetAddress+testAddress(chain)) {
			resp.StatusCode = http.StatusBadRequest
			resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"error": "Invalid address, decoded address is of unknown format"}`)))
			return resp, nil
		}
	}

	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

func TestClient_GetAddress(t *testing.T) {
	t.Parallel()

	t.Run("valid cases", func(t *testing.T) {

		var tests = []struct {
			chain        Blockchain
			address      string
			expectedInfo *AddressInfo
		}{
			{
				BSV, testAddress(BSV),
				&AddressInfo{
					Address:       testAddress(BSV),
					Balance:       "101556",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    174,
					TotalReceived: "66351012",
					TotalSent:     "66249456",
					TxIDs: []string{
						"12d9e1ed43444f03dee7381f66ecc206fe14719523a72d2c6536e1f4c4e05a14",
						"bc762632e1a02417ba0ee69b7d19059d3ba4f7905320f3930f22be6e9dde8610",
						"fcf4cda23a9394d446a421a15bfd7da0095fd701f2fc3c446d6a54e744caf1e2",
						"baee5d90e542b146eb3ce65e2b9d5f39c3a178b4134c583a6c9e379162c7604c",
						"b1e1d62e2ad74aab16741779bae5c03b64d35dba176687f88e475b9b5504fd6e",
					},
					Txs:                173661,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
			{
				BCH, testAddress(BCH),
				&AddressInfo{
					Address:       testAddress(BCH),
					Balance:       "3706237",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    104,
					TotalReceived: "1053381020454",
					TotalSent:     "1053377314217",
					TxIDs: []string{
						"fc2ee06ff4a22a630db1222ac1afa3595046854d03b12e6f423ee37e7db52be8",
						"640f5f203691199838efb0ab43a5c9f15cadbf4c573523058816f02590f95dff",
						"ac8e25de62f92dd2143c85b5487d2b77a34954c564bfe3047bc34639711c8b63",
					},
					Txs:                103219,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
			{
				BTC, testAddress(BTC),
				&AddressInfo{
					Address:       testAddress(BTC),
					Balance:       "4756408",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    1,
					TotalReceived: "209893262",
					TotalSent:     "205136854",
					TxIDs: []string{
						"bf8892afca454f5822dfe51eb8aebabab64b2803ea5ff8902b1bc6b0b1dd5e08",
						"e9da28382a504a1f86b063cc5b4c4d242af006ec61f4189310880da8f387ab3f",
						"3f696031a6e2a07364c5d5e6d829f3fd85fcf6f91f8b6fa7a4d0f24b563f3373",
					},
					Txs:                217,
					UnconfirmedBalance: "-854392",
					UnconfirmedTxs:     9,
				},
			},
			{
				BTG, testAddress(BTG),
				&AddressInfo{
					Address:       testAddress(BTG),
					Balance:       "121734541552",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    13,
					TotalReceived: "14394650955688",
					TotalSent:     "14272916414136",
					TxIDs: []string{
						"f52056973f853356d059e2cd05e214e8c405cab398240e811a54b622eff61c56",
						"ca5a3bbd6245598ec1af177b2d56e370139d96217600dec1455216ae87211314",
						"0ba336fe89ca26d04aaeabed8b2b252dba45004c389216af15290b6ff9179737",
					},
					Txs:                12287,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
			{
				DASH, testAddress(DASH),
				&AddressInfo{
					Address:       testAddress(DASH),
					Balance:       "0",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    1,
					TotalReceived: "54100000",
					TotalSent:     "54100000",
					TxIDs: []string{
						"7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e",
						"b9e4cebbf1cde3118ec62df72553a7d60fad78565028e8bf1c31edec23a1a393",
					},
					Txs:                2,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
			{
				DOGE, testAddress(DOGE),
				&AddressInfo{
					Address:       testAddress(DOGE),
					Balance:       "16765841411433",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    18,
					TotalReceived: "189519869773689377",
					TotalSent:     "189503103932277944",
					TxIDs: []string{
						"9a3bf978b4e48251f7c5c640afc6cf5cf9f0888b7d9cfca3b4db65d14b99bccf",
						"7b36a2713b5b3e12633a2a989dff177d31a0d49df50010e8822a5ce0ee70ab3d",
						"c65acb715eec3eec0dd405ff0da566ee6b5a7af3c2cfc87a2c4d87a7e1c1f01e",
					},
					Txs:                17885,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
			{
				LTC, testAddress(LTC),
				&AddressInfo{
					Address:       testAddress(LTC),
					Balance:       "4183327",
					ItemsOnPage:   1000,
					Page:          1,
					TotalPages:    1,
					TotalReceived: "97526951",
					TotalSent:     "93343624",
					TxIDs: []string{
						"72c9a6bc19049c6f455477af8561a726352a61f55c5a16f8165f09503951d94d",
						"1de19db94cab8f7bc4aaf15691c03732ed656ae8ce374e9b3d347520b563c49b",
						"2381fc53f03717267895a885109075f8ba275401fbea13f9e1ab8994c3017889",
					},
					Txs:                32,
					UnconfirmedBalance: "0",
					UnconfirmedTxs:     0,
				},
			},
		}

		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validAddressResponse{}))
		ctx := context.Background()

		for _, testCase := range tests {
			t.Run("chain "+testCase.chain.String()+": GetAddress("+testCase.address+")", func(t *testing.T) {
				info, err := c.GetAddress(ctx, testCase.chain, testCase.address)
				require.NoError(t, err)
				require.NotNil(t, info)

				// Check all fields and values
				assert.Equal(t, testCase.expectedInfo.Address, info.Address)
				assert.Equal(t, testCase.expectedInfo.Balance, info.Balance)
				assert.Equal(t, testCase.expectedInfo.ItemsOnPage, info.ItemsOnPage)
				assert.Equal(t, testCase.expectedInfo.Page, info.Page)
				assert.Equal(t, testCase.expectedInfo.TotalPages, info.TotalPages)
				assert.Equal(t, testCase.expectedInfo.TotalReceived, info.TotalReceived)
				assert.Equal(t, testCase.expectedInfo.TotalSent, info.TotalSent)
				assert.Equal(t, testCase.expectedInfo.TxIDs, info.TxIDs)
				assert.Equal(t, testCase.expectedInfo.Txs, info.Txs)
				assert.Equal(t, testCase.expectedInfo.UnconfirmedBalance, info.UnconfirmedBalance)
				assert.Equal(t, testCase.expectedInfo.UnconfirmedTxs, info.UnconfirmedTxs)
			})
		}
	})

	t.Run("missing address", func(t *testing.T) {
		type testData struct {
			chain   Blockchain
			address string
			err     error
		}
		var testCases []testData
		for _, chain := range getAddressBlockchains {
			testCases = append(testCases, testData{chain: chain, address: "", err: ErrInvalidAddress})
			testCases = append(testCases, testData{chain: chain, address: "12345", err: ErrInvalidAddress})
			testCases = append(testCases, testData{chain: chain, address: "invalid-tx-hex", err: ErrInvalidAddress})
		}

		c := NewClient(WithHTTPClient(&validAddressResponse{}))
		ctx := context.Background()

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": missing-address", func(t *testing.T) {
				info, err := c.GetAddress(ctx, testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
				assert.ErrorIs(t, err, testCase.err)
			})
		}
	})

	t.Run("unsupported chain", func(t *testing.T) {
		c := NewClient(WithHTTPClient(&validAddressResponse{}))
		ctx := context.Background()
		info, err := c.GetAddress(ctx, ETH, testAddress(ETH))
		require.Error(t, err)
		require.Nil(t, info)
		assert.ErrorIs(t, err, ErrUnsupportedBlockchain)
	})

	t.Run("error cases", func(t *testing.T) {

		type testData struct {
			chain   Blockchain
			address string
		}
		var testCases []testData
		for _, chain := range getAddressBlockchains {
			testCases = append(testCases, testData{chain: chain, address: testAddress(chain)})
		}

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": address invalid", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorInvalidAddress{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": http req error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqErr{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": missing body contents", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqNoBodyErr{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": error with resp", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqWithRespErr{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": invalid json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadJSONResponse{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": invalid error json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadErrorJSONResponse{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": missing api key", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorMissingAPIKey{}))
				info, err := c.GetAddress(context.Background(), testCase.chain, testCase.address)
				require.Error(t, err)
				require.Nil(t, info)
			})
		}
	})
}

func ExampleClient_GetAddress() {
	c := NewClient(WithHTTPClient(&validAddressResponse{}))
	info, _ := c.GetAddress(context.Background(), BSV, testAddress(BSV))
	fmt.Println("address found: " + info.Address)
	// Output:address found: 1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL
}

func BenchmarkClient_GetAddress(b *testing.B) {
	c := NewClient(WithHTTPClient(&validTxResponse{}))
	ctx := context.Background()
	address := testAddress(BSV)
	for i := 0; i < b.N; i++ {
		_, _ = c.GetAddress(ctx, BSV, address)
	}
}
