package nownodes

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// validTxResponse will return a valid tx for all supported blockchains
type validTxResponse struct{}

func (v *validTxResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	// Transaction data
	transactionResponses := map[string]string{
		BCH.String():  `{"txid":"fb91f4b1969c7dcff3a1199b26ba9b33af658cec98641a32740f06f0e09b0efe","version":2,"vin":[{"txid":"2c18adefd40022b3417d959fcff893cceba5e69469aa7f4cf1a002898c626e69","vout":1,"n":0,"addresses":["bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn"],"isAddress":true,"value":"4162642","hex":"414d942623618a94335de26cf9437e5f071029ee4749b9ed5be422c49cd9f158dfe170bdc95b5ec2d2656671acf9d33a0b8486a1c6ca4fd367e5a1cc7b66483100412102e1ee329e4bce33ee828320743b261ff59102e83e36e35c7875910a1ea78b0508"}],"vout":[{"value":"340976","n":0,"hex":"76a9143c282e546e6bae90873ed2bc51aa29ba4aac9eb488ac","addresses":["bitcoincash:qq7zstj5de46ayy88mftc5d29xay4ty7ksfm6h9237"],"isAddress":true},{"value":"3821446","n":1,"hex":"76a91490258f19881785794d4a522c3fe28e5ea2599c9988ac","addresses":["bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn"],"isAddress":true}],"blockHash":"0000000000000000007edfce162d75a522d8e6b38745b5e9b53d2ff05ccd1faf","blockHeight":725003,"confirmations":1,"blockTime":1643485950,"value":"4162422","valueIn":"4162642","fees":"220","hex":"0200000001696e628c8902a0f14c7faa6994e6a5ebcc93f8cf9f957d41b32200d4efad182c0100000064414d942623618a94335de26cf9437e5f071029ee4749b9ed5be422c49cd9f158dfe170bdc95b5ec2d2656671acf9d33a0b8486a1c6ca4fd367e5a1cc7b66483100412102e1ee329e4bce33ee828320743b261ff59102e83e36e35c7875910a1ea78b05080000000002f0330500000000001976a9143c282e546e6bae90873ed2bc51aa29ba4aac9eb488ac864f3a00000000001976a91490258f19881785794d4a522c3fe28e5ea2599c9988ac00000000"}`,
		BSV.String():  `{"txid":"17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd","version":1,"vin":[{"txid":"cab4b07235120ed66aabcfc907b42be6c8782418461aebabd2ba58c08ca38ccc","vout":2,"sequence":4294967295,"n":0,"addresses":["1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL"],"isAddress":true,"value":"546","hex":"483045022100977a4cbf4f34efc54ff56a1d4b74148836e86b045ad75fd4c27729e2f3c9cf9e02205ed85454fdecf4345fdd1bf19a64db3268c3fe00d72615de2a1a24c92defbde5412102cfbb8f465fa014012bd44407974fbd13f239b9b5e9586db75191acaa642336f0"}],"vout":[{"value":"0","n":0,"hex":"006a0372756e0105036679784ca67b22696e223a312c22726566223a5b5d2c226f7574223a5b5d2c2264656c223a5b2265353066616364323332663663333037326337313538393333373839313437376432343037373235393839636161376439363062383662303533633736323366225d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c2264657374726f79222c5b5d5d7d5d7d","addresses":[],"isAddress":false},{"value":"450","n":1,"hex":"76a914f08d4568df6be038700227e70105b251455abaf188ac","addresses":["1NvvQjKN4GsyA9Y2kUT8PRvocAPJgCneFZ"],"isAddress":true}],"blockHash":"00000000000000000a032702d724591574cae47e729acdaf8b8a990adda3f72e","blockHeight":723772,"confirmations":622,"blockTime":1643111792,"value":"450","valueIn":"546","fees":"96","hex":"0100000001cc8ca38cc058bad2abeb1a46182478c8e62bb407c9cfab6ad60e123572b0b4ca020000006b483045022100977a4cbf4f34efc54ff56a1d4b74148836e86b045ad75fd4c27729e2f3c9cf9e02205ed85454fdecf4345fdd1bf19a64db3268c3fe00d72615de2a1a24c92defbde5412102cfbb8f465fa014012bd44407974fbd13f239b9b5e9586db75191acaa642336f0ffffffff020000000000000000b4006a0372756e0105036679784ca67b22696e223a312c22726566223a5b5d2c226f7574223a5b5d2c2264656c223a5b2265353066616364323332663663333037326337313538393333373839313437376432343037373235393839636161376439363062383662303533633736323366225d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c2264657374726f79222c5b5d5d7d5d7dc2010000000000001976a914f08d4568df6be038700227e70105b251455abaf188ac00000000"}`,
		BTC.String():  `{"txid":"050ead9fbd6360771541c734fa2324e5caa13f62a9598a39ca730aeb19e8c89a","version":1,"vin":[{"txid":"67017bbf2023a94140b2091e7a413a027c9565f9f150d8cfd44c12af3286ed4f","vout":1,"sequence":4294967295,"n":0,"addresses":["17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"],"isAddress":true,"value":"96760","hex":"4730440220144776296a112aab37729e04e93725d7cee443b673267b991953950f5c00920f022010e31ad967b9bb7182650a71b527a46578f776518056d8767242b7f1c0c4f9350121023fb7b226303b63e5caf6e93e79aa28b0b4f5f1b9e48ca97af62277cccbc5127e"}],"vout":[{"value":"0","n":0,"hex":"6a36c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185c","addresses":["OP_RETURN c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185c"],"isAddress":false},{"value":"96496","n":1,"hex":"76a91448dfc8dbdd463b27ba60fe6da4f8751199f44a5388ac","addresses":["17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"],"isAddress":true}],"blockHash":"00000000000000000008674e0259616fe31ca686ef6dcbd0ec60636713fe910d","blockHeight":720943,"confirmations":1,"blockTime":1643486938,"value":"96496","valueIn":"96760","fees":"264","hex":"01000000014fed8632af124cd4cfd850f1f965957c023a417a1e09b24041a92320bf7b0167010000006a4730440220144776296a112aab37729e04e93725d7cee443b673267b991953950f5c00920f022010e31ad967b9bb7182650a71b527a46578f776518056d8767242b7f1c0c4f9350121023fb7b226303b63e5caf6e93e79aa28b0b4f5f1b9e48ca97af62277cccbc5127effffffff020000000000000000386a36c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185cf0780100000000001976a91448dfc8dbdd463b27ba60fe6da4f8751199f44a5388ac00000000"}`,
		BTG.String():  `{"txid":"934989d8e6e1fe9bc3d7508479df85e4757e6f0ceee613e7f8736e4e6b344a4a","version":1,"vin":[{"txid":"4ac8d33e95c3944428ede85c68912951f6097022974ff4b7cf7e274dbf534686","sequence":4294967295,"n":0,"addresses":["ATTav2PtmotBZwxgWjrZCgaZpE89kcJ29B"],"isAddress":true,"value":"759789817","hex":"1600143c0829804f4c55122d8f403f614ce4f845290fdb"}],"vout":[{"value":"14276945","n":0,"hex":"76a9140cb60a52559620e5de9a297612d49f55f7fd14ea88ac","addresses":["GK18bp4UzC6wqYKKNLkaJ3hzQazTc3TWBw"],"isAddress":true},{"value":"745512674","n":1,"hex":"a9148bcd7f6402f5fd50f34850e2c5f2e45e4c1702c887","addresses":["AUX5kPSTQeosDXmTroBZPLHv7NNXZYZkvX"],"isAddress":true}],"blockHash":"0000000174db2a8a13531cd8a7f42a3a2eca5c3e4b2a909a196dcb5b9dc382d1","blockHeight":723261,"confirmations":11,"blockTime":1643483952,"value":"759789619","valueIn":"759789817","fees":"198","hex":"01000000000101864653bf4d277ecfb7f44f97227009f6512991685ce8ed284494c3953ed3c84a00000000171600143c0829804f4c55122d8f403f614ce4f845290fdbffffffff0251d9d900000000001976a9140cb60a52559620e5de9a297612d49f55f7fd14ea88ace29e6f2c0000000017a9148bcd7f6402f5fd50f34850e2c5f2e45e4c1702c88702483045022100b23f6fbadf3c4b22ccaa7245cb8c1cb4340f32667ee4069a1a843f7681a24d080220199f827e9701b281f62cddb42df6de35406c82b2c2e7f3b4cdaf0034c7b23fef412103d0c56dd160c29607cf4463d619822946666f9998b2ae86b54a5821cab704f2b300000000"}`,
		DASH.String(): `{"txid":"7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e","version":2,"lockTime":1613396,"vin":[{"txid":"b9e4cebbf1cde3118ec62df72553a7d60fad78565028e8bf1c31edec23a1a393","sequence":4294967294,"n":0,"addresses":["Xe7tPVUvDpt52h2KykMpj4hmh8VsAaPCgt"],"isAddress":true,"value":"54100000","hex":"4830450221009ceb5a9f743de29353e077ef642b4a741d88d60e8737918ed7e060b6017750af022062e5b15f4ac3f9f5ea97124d3ce3e351a950f7018adb22d28b3a0bb86594998f0121036fdd18e0e1ff3989431beac0aef1a5e75f51d745ce4a54afd8a3a1968592275d"}],"vout":[{"value":"54099776","n":0,"hex":"a914581cbcc7c2a93077d836c232130cfbcf3987f0ce87","addresses":["7aSYeL7uF9HtxVYiTX8Ew6wFYkcE3veAqj"],"isAddress":true}],"blockHash":"000000000000001e507180f6ab9aa1d0541d562592af4e5c77a60427c4e174e9","blockHeight":1613398,"confirmations":7,"blockTime":1643487799,"value":"54099776","valueIn":"54100000","fees":"224","hex":"020000000193a3a123eced311cbfe828505678ad0fd6a75325f72dc68e11e3cdf1bbcee4b9000000006b4830450221009ceb5a9f743de29353e077ef642b4a741d88d60e8737918ed7e060b6017750af022062e5b15f4ac3f9f5ea97124d3ce3e351a950f7018adb22d28b3a0bb86594998f0121036fdd18e0e1ff3989431beac0aef1a5e75f51d745ce4a54afd8a3a1968592275dfeffffff01407f39030000000017a914581cbcc7c2a93077d836c232130cfbcf3987f0ce87549e1800"}`,
		DOGE.String(): `{"txid":"6b22cc41b1206b6f39568bca5ca9e32ca3e0f6f4e0a68e2b126913e7d6620543","version":1,"vin":[{"txid":"44e93228de0618bd425e30657485ac0eebb56e205905698c9a139bdb0092e964","vout":1,"sequence":4294967295,"n":0,"addresses":["ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a"],"isAddress":true,"value":"6894572842640","hex":"004730440220646f43d6d57b850206a8dde6750cd3a6ba4f77cd26452505757a49f0345fd547022057d093128e611a2d1c732058f38a4528dffaf13bf8bf221f03c17165afdd79360147304402203f49525951952b192559767c7b6228a24527657786dae173aab92d883405f56a02207f70dec8ea9a15078ea11f6737c21be45e5059e421e53035bd1390f3290661920147522102adf2cb5afd730171a425d263014e9307d71edc352b4c3c9b750eecdc95ef70d021027105672d0cf8269ca3ea757754049eeec8da3a7e963d2786f06547cd78c28be252ae"}],"vout":[{"value":"88041747863","n":0,"spent":true,"hex":"76a91473d7fc810d7d02988219702c5296eed5e2f9449988ac","addresses":["DFhczK7w4gjGrNjFTgLFEYbL8Zs2YQA1dQ"],"isAddress":true},{"value":"6806530086777","n":1,"hex":"a914db72653436f25884f2ab2bf050d06e805907dd0e87","addresses":["ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a"],"isAddress":true}],"blockHash":"cf67498603ab0c8dd66324688dc69d1e4d3edd4e894ea993cbf2931659d59b36","blockHeight":4082794,"confirmations":9,"blockTime":1643487708,"value":"6894571834640","valueIn":"6894572842640","fees":"1008000","hex":"010000000164e99200db9b139a8c690559206eb5eb0eac857465305e42bd1806de2832e94401000000d9004730440220646f43d6d57b850206a8dde6750cd3a6ba4f77cd26452505757a49f0345fd547022057d093128e611a2d1c732058f38a4528dffaf13bf8bf221f03c17165afdd79360147304402203f49525951952b192559767c7b6228a24527657786dae173aab92d883405f56a02207f70dec8ea9a15078ea11f6737c21be45e5059e421e53035bd1390f3290661920147522102adf2cb5afd730171a425d263014e9307d71edc352b4c3c9b750eecdc95ef70d021027105672d0cf8269ca3ea757754049eeec8da3a7e963d2786f06547cd78c28be252aeffffffff029775b27f140000001976a91473d7fc810d7d02988219702c5296eed5e2f9449988ac79d7cec43006000017a914db72653436f25884f2ab2bf050d06e805907dd0e8700000000"}`,
		LTC.String():  `{"txid":"dfca839b7686a458e94001e53df4bd3bbe967d7ba5622f67b38cd6e2650bb37a","version":1,"vin":[{"txid":"95f0a38d47ea3fd21d467a683ef744391d777000ac1dec96f2bd24759cbd76f5","vout":1,"sequence":4294967295,"n":0,"addresses":["ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps"],"isAddress":true,"value":"4890906"}],"vout":[{"value":"100000","n":0,"hex":"a914777762c97ceb6cd2ec0eded08868bd953e838f7987","addresses":["MJnqfLNiC3LvH2efxjP4yNjdKbVgpGf3Ar"],"isAddress":true},{"value":"4785482","n":1,"spent":true,"hex":"0014d1ea11d9f10744ae033327ecbe8282283088b45e","addresses":["ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps"],"isAddress":true}],"blockHash":"b8f2cb74105dd4e7b8850bc1743cf3174f7365ef5bce6f015a83e8918e2ad58f","blockHeight":2202057,"confirmations":7,"blockTime":1643487498,"value":"4885482","valueIn":"4890906","fees":"5424","hex":"01000000000101f576bd9c7524bdf296ec1dac0070771d3944f73e687a461dd23fea478da3f0950100000000ffffffff02a08601000000000017a914777762c97ceb6cd2ec0eded08868bd953e838f79874a05490000000000160014d1ea11d9f10744ae033327ecbe8282283088b45e02483045022100fd62f1f896e0ec3d75e84c977f91d163fe28bc576827bd39015ef507ea4bb3ee02201ff0f5baa4a2336bbb9bc92029d10dfa82db636b5ca178d3a7b621dfefdd8ac4012102c4303428959c2d86c3c742586651f384685b99ee007a186e7a1326f5548c682e00000000"}`,
	}

	// Valid response (get tx)
	for _, chain := range getTransactionBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeGetTx+testTxID(chain)) {
			resp.StatusCode = http.StatusOK
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(transactionResponses[chain.String()])))
			return resp, nil
		}
	}

	// Valid response (send tx)
	for _, chain := range sendTransactionBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeSendTx+testTxHex(chain)) {
			resp.StatusCode = http.StatusOK
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"result":"` + testTxHexID(chain) + `"}`)))
			return resp, nil
		}
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

// errorTxNotFoundResponse will return an error for the tx response
type errorTxNotFoundResponse struct{}

func (v *errorTxNotFoundResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	// Error response (get tx)
	for _, chain := range getTransactionBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeGetTx+testTxID(chain)) {
			resp.StatusCode = http.StatusBadRequest
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error": "Transaction '` + testTxID(chain) + `' not found"}`)))
			return resp, nil
		}
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

// errorSendTxErrorResponse will return an error for the "send tx" response
type errorSendTxErrorResponse struct{}

func (v *errorSendTxErrorResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	// Error response (send tx)
	for _, chain := range sendTransactionBlockchains {
		if strings.Contains(req.Host, chain.BlockBookURL()) && strings.Contains(req.URL.String(), routeSendTx+testTxHex(chain)) {
			resp.StatusCode = http.StatusBadRequest
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error": "-27: Transaction already in the mempool"}`)))
			return resp, nil
		}
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

func TestClient_GetTransaction(t *testing.T) {
	t.Parallel()

	t.Run("valid cases", func(t *testing.T) {

		var tests = []struct {
			chain        Blockchain
			txID         string
			expectedInfo *TransactionInfo
		}{
			{
				BSV, testTxID(BSV),
				&TransactionInfo{
					BlockHash:     "00000000000000000a032702d724591574cae47e729acdaf8b8a990adda3f72e",
					BlockHeight:   723772,
					BlockTime:     1643111792,
					Confirmations: 622,
					Fees:          "96",
					Hex:           "0100000001cc8ca38cc058bad2abeb1a46182478c8e62bb407c9cfab6ad60e123572b0b4ca020000006b483045022100977a4cbf4f34efc54ff56a1d4b74148836e86b045ad75fd4c27729e2f3c9cf9e02205ed85454fdecf4345fdd1bf19a64db3268c3fe00d72615de2a1a24c92defbde5412102cfbb8f465fa014012bd44407974fbd13f239b9b5e9586db75191acaa642336f0ffffffff020000000000000000b4006a0372756e0105036679784ca67b22696e223a312c22726566223a5b5d2c226f7574223a5b5d2c2264656c223a5b2265353066616364323332663663333037326337313538393333373839313437376432343037373235393839636161376439363062383662303533633736323366225d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c2264657374726f79222c5b5d5d7d5d7dc2010000000000001976a914f08d4568df6be038700227e70105b251455abaf188ac00000000",
					TxID:          testTxID(BSV),
					Value:         "450",
					ValueIn:       "546",
					Version:       1,
					Vin: []*Input{{
						Addresses: []string{"1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL"},
						Hex:       "483045022100977a4cbf4f34efc54ff56a1d4b74148836e86b045ad75fd4c27729e2f3c9cf9e02205ed85454fdecf4345fdd1bf19a64db3268c3fe00d72615de2a1a24c92defbde5412102cfbb8f465fa014012bd44407974fbd13f239b9b5e9586db75191acaa642336f0",
						IsAddress: true,
						N:         0,
						Sequence:  4294967295,
						TxID:      "cab4b07235120ed66aabcfc907b42be6c8782418461aebabd2ba58c08ca38ccc",
						Value:     "546",
						VOut:      2,
					}},
					VOut: []*Output{
						{
							Hex:       "006a0372756e0105036679784ca67b22696e223a312c22726566223a5b5d2c226f7574223a5b5d2c2264656c223a5b2265353066616364323332663663333037326337313538393333373839313437376432343037373235393839636161376439363062383662303533633736323366225d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c2264657374726f79222c5b5d5d7d5d7d",
							IsAddress: false,
							N:         0,
							Value:     "0",
						},
						{
							Addresses: []string{"1NvvQjKN4GsyA9Y2kUT8PRvocAPJgCneFZ"},
							Hex:       "76a914f08d4568df6be038700227e70105b251455abaf188ac",
							IsAddress: true,
							N:         1,
							Value:     "450",
						},
					},
				},
			},
			{
				BCH, testTxID(BCH),
				&TransactionInfo{
					BlockHash:     "0000000000000000007edfce162d75a522d8e6b38745b5e9b53d2ff05ccd1faf",
					BlockHeight:   725003,
					BlockTime:     1643485950,
					Confirmations: 1,
					Fees:          "220",
					Hex:           "0200000001696e628c8902a0f14c7faa6994e6a5ebcc93f8cf9f957d41b32200d4efad182c0100000064414d942623618a94335de26cf9437e5f071029ee4749b9ed5be422c49cd9f158dfe170bdc95b5ec2d2656671acf9d33a0b8486a1c6ca4fd367e5a1cc7b66483100412102e1ee329e4bce33ee828320743b261ff59102e83e36e35c7875910a1ea78b05080000000002f0330500000000001976a9143c282e546e6bae90873ed2bc51aa29ba4aac9eb488ac864f3a00000000001976a91490258f19881785794d4a522c3fe28e5ea2599c9988ac00000000",
					TxID:          testBCHTxID,
					Value:         "4162422",
					ValueIn:       "4162642",
					Version:       2,
					Vin: []*Input{{
						Addresses: []string{"bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn"},
						Hex:       "414d942623618a94335de26cf9437e5f071029ee4749b9ed5be422c49cd9f158dfe170bdc95b5ec2d2656671acf9d33a0b8486a1c6ca4fd367e5a1cc7b66483100412102e1ee329e4bce33ee828320743b261ff59102e83e36e35c7875910a1ea78b0508",
						IsAddress: true,
						N:         0,
						TxID:      "2c18adefd40022b3417d959fcff893cceba5e69469aa7f4cf1a002898c626e69",
						Value:     "4162642",
						VOut:      1,
					}},
					VOut: []*Output{{
						Addresses: []string{"bitcoincash:qq7zstj5de46ayy88mftc5d29xay4ty7ksfm6h9237"},
						Hex:       "76a9143c282e546e6bae90873ed2bc51aa29ba4aac9eb488ac",
						IsAddress: true,
						N:         0,
						Value:     "340976",
					}, {
						Addresses: []string{"bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn"},
						Hex:       "76a91490258f19881785794d4a522c3fe28e5ea2599c9988ac",
						IsAddress: true,
						N:         1,
						Value:     "3821446",
					}},
				},
			},
			{
				BTC, testTxID(BTC),
				&TransactionInfo{
					BlockHash:     "00000000000000000008674e0259616fe31ca686ef6dcbd0ec60636713fe910d",
					BlockHeight:   720943,
					BlockTime:     1643486938,
					Confirmations: 1,
					Fees:          "264",
					Hex:           "01000000014fed8632af124cd4cfd850f1f965957c023a417a1e09b24041a92320bf7b0167010000006a4730440220144776296a112aab37729e04e93725d7cee443b673267b991953950f5c00920f022010e31ad967b9bb7182650a71b527a46578f776518056d8767242b7f1c0c4f9350121023fb7b226303b63e5caf6e93e79aa28b0b4f5f1b9e48ca97af62277cccbc5127effffffff020000000000000000386a36c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185cf0780100000000001976a91448dfc8dbdd463b27ba60fe6da4f8751199f44a5388ac00000000",
					TxID:          testBTCTxID,
					Value:         "96496",
					ValueIn:       "96760",
					Version:       1,
					Vin: []*Input{{
						Addresses: []string{"17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"},
						Hex:       "4730440220144776296a112aab37729e04e93725d7cee443b673267b991953950f5c00920f022010e31ad967b9bb7182650a71b527a46578f776518056d8767242b7f1c0c4f9350121023fb7b226303b63e5caf6e93e79aa28b0b4f5f1b9e48ca97af62277cccbc5127e",
						IsAddress: true,
						N:         0,
						Sequence:  4294967295,
						TxID:      "67017bbf2023a94140b2091e7a413a027c9565f9f150d8cfd44c12af3286ed4f",
						Value:     "96760",
					}},
					VOut: []*Output{{
						Addresses: []string{
							"OP_RETURN c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185c",
						},
						Hex:       "6a36c6329c9c6c21e8d9d1392af3314ddcdba5f67f95cfd387d6970349929817341b8b775174b3068836ff73cf23a9e3beef1b45c96f185c",
						IsAddress: false,
						N:         0,
						Value:     "0",
					}, {
						Addresses: []string{"17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"},
						Hex:       "76a91448dfc8dbdd463b27ba60fe6da4f8751199f44a5388ac",
						IsAddress: true,
						N:         1,
						Value:     "96496",
					}},
				},
			},
			{
				BTG, testTxID(BTG),
				&TransactionInfo{
					BlockHash:     "0000000174db2a8a13531cd8a7f42a3a2eca5c3e4b2a909a196dcb5b9dc382d1",
					BlockHeight:   723261,
					BlockTime:     1643483952,
					Confirmations: 11,
					Fees:          "198",
					Hex:           "01000000000101864653bf4d277ecfb7f44f97227009f6512991685ce8ed284494c3953ed3c84a00000000171600143c0829804f4c55122d8f403f614ce4f845290fdbffffffff0251d9d900000000001976a9140cb60a52559620e5de9a297612d49f55f7fd14ea88ace29e6f2c0000000017a9148bcd7f6402f5fd50f34850e2c5f2e45e4c1702c88702483045022100b23f6fbadf3c4b22ccaa7245cb8c1cb4340f32667ee4069a1a843f7681a24d080220199f827e9701b281f62cddb42df6de35406c82b2c2e7f3b4cdaf0034c7b23fef412103d0c56dd160c29607cf4463d619822946666f9998b2ae86b54a5821cab704f2b300000000",
					TxID:          testBTGTxID,
					Value:         "759789619",
					ValueIn:       "759789817",
					Version:       1,
					Vin: []*Input{{
						Addresses: []string{"ATTav2PtmotBZwxgWjrZCgaZpE89kcJ29B"},
						Hex:       "1600143c0829804f4c55122d8f403f614ce4f845290fdb",
						IsAddress: true,
						N:         0,
						Sequence:  4294967295,
						TxID:      "4ac8d33e95c3944428ede85c68912951f6097022974ff4b7cf7e274dbf534686",
						Value:     "759789817",
					}},
					VOut: []*Output{{
						Addresses: []string{
							"GK18bp4UzC6wqYKKNLkaJ3hzQazTc3TWBw",
						},
						Hex:       "76a9140cb60a52559620e5de9a297612d49f55f7fd14ea88ac",
						IsAddress: true,
						N:         0,
						Value:     "14276945",
					}, {
						Addresses: []string{"AUX5kPSTQeosDXmTroBZPLHv7NNXZYZkvX"},
						Hex:       "a9148bcd7f6402f5fd50f34850e2c5f2e45e4c1702c887",
						IsAddress: true,
						N:         1,
						Value:     "745512674",
					}},
				},
			},
			{
				DASH, testTxID(DASH),
				&TransactionInfo{
					BlockHash:     "000000000000001e507180f6ab9aa1d0541d562592af4e5c77a60427c4e174e9",
					BlockHeight:   1613398,
					BlockTime:     1643487799,
					Confirmations: 7,
					LockTime:      1613396,
					Fees:          "224",
					Hex:           "020000000193a3a123eced311cbfe828505678ad0fd6a75325f72dc68e11e3cdf1bbcee4b9000000006b4830450221009ceb5a9f743de29353e077ef642b4a741d88d60e8737918ed7e060b6017750af022062e5b15f4ac3f9f5ea97124d3ce3e351a950f7018adb22d28b3a0bb86594998f0121036fdd18e0e1ff3989431beac0aef1a5e75f51d745ce4a54afd8a3a1968592275dfeffffff01407f39030000000017a914581cbcc7c2a93077d836c232130cfbcf3987f0ce87549e1800",
					TxID:          testDASHTxID,
					Value:         "54099776",
					ValueIn:       "54100000",
					Version:       2,
					Vin: []*Input{{
						Addresses: []string{"Xe7tPVUvDpt52h2KykMpj4hmh8VsAaPCgt"},
						Hex:       "4830450221009ceb5a9f743de29353e077ef642b4a741d88d60e8737918ed7e060b6017750af022062e5b15f4ac3f9f5ea97124d3ce3e351a950f7018adb22d28b3a0bb86594998f0121036fdd18e0e1ff3989431beac0aef1a5e75f51d745ce4a54afd8a3a1968592275d",
						IsAddress: true,
						N:         0,
						Sequence:  4294967294,
						TxID:      "b9e4cebbf1cde3118ec62df72553a7d60fad78565028e8bf1c31edec23a1a393",
						Value:     "54100000",
					}},
					VOut: []*Output{{
						Addresses: []string{
							"7aSYeL7uF9HtxVYiTX8Ew6wFYkcE3veAqj",
						},
						Hex:       "a914581cbcc7c2a93077d836c232130cfbcf3987f0ce87",
						IsAddress: true,
						N:         0,
						Value:     "54099776",
					}},
				},
			},
			{
				DOGE, testTxID(DOGE),
				&TransactionInfo{
					BlockHash:     "cf67498603ab0c8dd66324688dc69d1e4d3edd4e894ea993cbf2931659d59b36",
					BlockHeight:   4082794,
					BlockTime:     1643487708,
					Confirmations: 9,
					Fees:          "1008000",
					Hex:           "010000000164e99200db9b139a8c690559206eb5eb0eac857465305e42bd1806de2832e94401000000d9004730440220646f43d6d57b850206a8dde6750cd3a6ba4f77cd26452505757a49f0345fd547022057d093128e611a2d1c732058f38a4528dffaf13bf8bf221f03c17165afdd79360147304402203f49525951952b192559767c7b6228a24527657786dae173aab92d883405f56a02207f70dec8ea9a15078ea11f6737c21be45e5059e421e53035bd1390f3290661920147522102adf2cb5afd730171a425d263014e9307d71edc352b4c3c9b750eecdc95ef70d021027105672d0cf8269ca3ea757754049eeec8da3a7e963d2786f06547cd78c28be252aeffffffff029775b27f140000001976a91473d7fc810d7d02988219702c5296eed5e2f9449988ac79d7cec43006000017a914db72653436f25884f2ab2bf050d06e805907dd0e8700000000",
					TxID:          testDOGETxID,
					Value:         "6894571834640",
					ValueIn:       "6894572842640",
					Version:       1,
					Vin: []*Input{{
						Addresses: []string{"ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a"},
						Hex:       "004730440220646f43d6d57b850206a8dde6750cd3a6ba4f77cd26452505757a49f0345fd547022057d093128e611a2d1c732058f38a4528dffaf13bf8bf221f03c17165afdd79360147304402203f49525951952b192559767c7b6228a24527657786dae173aab92d883405f56a02207f70dec8ea9a15078ea11f6737c21be45e5059e421e53035bd1390f3290661920147522102adf2cb5afd730171a425d263014e9307d71edc352b4c3c9b750eecdc95ef70d021027105672d0cf8269ca3ea757754049eeec8da3a7e963d2786f06547cd78c28be252ae",
						IsAddress: true,
						N:         0,
						Sequence:  4294967295,
						TxID:      "44e93228de0618bd425e30657485ac0eebb56e205905698c9a139bdb0092e964",
						Value:     "6894572842640",
						VOut:      1,
					}},
					VOut: []*Output{{
						Addresses: []string{
							"DFhczK7w4gjGrNjFTgLFEYbL8Zs2YQA1dQ",
						},
						Hex:       "76a91473d7fc810d7d02988219702c5296eed5e2f9449988ac",
						IsAddress: true,
						Spent:     true,
						N:         0,
						Value:     "88041747863",
					}, {
						Addresses: []string{
							"ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a",
						},
						Hex:       "a914db72653436f25884f2ab2bf050d06e805907dd0e87",
						IsAddress: true,
						N:         1,
						Value:     "6806530086777",
					}},
				},
			},
			{
				LTC, testTxID(LTC),
				&TransactionInfo{
					BlockHash:     "b8f2cb74105dd4e7b8850bc1743cf3174f7365ef5bce6f015a83e8918e2ad58f",
					BlockHeight:   2202057,
					BlockTime:     1643487498,
					Confirmations: 7,
					Fees:          "5424",
					Hex:           "01000000000101f576bd9c7524bdf296ec1dac0070771d3944f73e687a461dd23fea478da3f0950100000000ffffffff02a08601000000000017a914777762c97ceb6cd2ec0eded08868bd953e838f79874a05490000000000160014d1ea11d9f10744ae033327ecbe8282283088b45e02483045022100fd62f1f896e0ec3d75e84c977f91d163fe28bc576827bd39015ef507ea4bb3ee02201ff0f5baa4a2336bbb9bc92029d10dfa82db636b5ca178d3a7b621dfefdd8ac4012102c4303428959c2d86c3c742586651f384685b99ee007a186e7a1326f5548c682e00000000",
					TxID:          testLTCTxID,
					Value:         "4885482",
					ValueIn:       "4890906",
					Version:       1,
					Vin: []*Input{{
						Addresses: []string{"ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps"},
						Hex:       "004730440220646f43d6d57b850206a8dde6750cd3a6ba4f77cd26452505757a49f0345fd547022057d093128e611a2d1c732058f38a4528dffaf13bf8bf221f03c17165afdd79360147304402203f49525951952b192559767c7b6228a24527657786dae173aab92d883405f56a02207f70dec8ea9a15078ea11f6737c21be45e5059e421e53035bd1390f3290661920147522102adf2cb5afd730171a425d263014e9307d71edc352b4c3c9b750eecdc95ef70d021027105672d0cf8269ca3ea757754049eeec8da3a7e963d2786f06547cd78c28be252ae",
						IsAddress: true,
						N:         0,
						Sequence:  4294967295,
						TxID:      "95f0a38d47ea3fd21d467a683ef744391d777000ac1dec96f2bd24759cbd76f5",
						Value:     "4890906",
						VOut:      1,
					}},
					VOut: []*Output{{
						Addresses: []string{
							"MJnqfLNiC3LvH2efxjP4yNjdKbVgpGf3Ar",
						},
						Hex:       "a914777762c97ceb6cd2ec0eded08868bd953e838f7987",
						IsAddress: true,
						N:         0,
						Value:     "100000",
					}, {
						Addresses: []string{
							"ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps",
						},
						Hex:       "0014d1ea11d9f10744ae033327ecbe8282283088b45e",
						IsAddress: true,
						Spent:     true,
						N:         1,
						Value:     "4785482",
					}},
				},
			},
		}

		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()

		for _, testCase := range tests {
			t.Run("chain "+testCase.chain.String()+": GetTransaction("+testCase.txID+")", func(t *testing.T) {
				info, err := c.GetTransaction(ctx, testCase.chain, testCase.txID)
				require.NoError(t, err)
				require.NotNil(t, info)

				// Check all fields and values
				assert.Equal(t, testCase.expectedInfo.BlockHash, info.BlockHash)
				assert.Equal(t, testCase.expectedInfo.BlockHeight, info.BlockHeight)
				assert.Equal(t, testCase.expectedInfo.BlockTime, info.BlockTime)
				assert.Equal(t, testCase.expectedInfo.Confirmations, info.Confirmations)
				assert.Equal(t, testCase.expectedInfo.Fees, info.Fees)
				assert.Equal(t, testCase.expectedInfo.Hex, info.Hex)
				assert.Equal(t, testCase.expectedInfo.TxID, info.TxID)
				assert.Equal(t, testCase.expectedInfo.Value, info.Value)
				assert.Equal(t, testCase.expectedInfo.ValueIn, info.ValueIn)
				assert.Equal(t, testCase.expectedInfo.Version, info.Version)

				// Check inputs and outputs
				assert.Equal(t, len(testCase.expectedInfo.Vin), len(info.Vin))
				assert.Equal(t, len(testCase.expectedInfo.VOut), len(info.VOut))
				// todo: check each input and output value
			})
		}
	})

	t.Run("missing or invalid tx id", func(t *testing.T) {
		type testData struct {
			chain Blockchain
			txID  string
			err   error
		}
		var testCases []testData
		for _, chain := range getTransactionBlockchains {
			testCases = append(testCases, testData{chain: chain, txID: "", err: ErrInvalidTxID})
			testCases = append(testCases, testData{chain: chain, txID: "12345", err: ErrInvalidTxID})
			testCases = append(testCases, testData{chain: chain, txID: "invalid-tx-hex", err: ErrInvalidTxID})
		}

		c := NewClient(WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": missing-tx", func(t *testing.T) {
				info, err := c.GetTransaction(ctx, testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
				assert.ErrorIs(t, err, testCase.err)
			})
		}
	})

	t.Run("unsupported chain", func(t *testing.T) {
		c := NewClient(WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()
		getTransactionBlockchains = []Blockchain{BSV}
		info, err := c.GetTransaction(ctx, BTC, testTxID(BTC))
		require.Error(t, err)
		require.Nil(t, info)
		assert.ErrorIs(t, err, ErrUnsupportedBlockchain)
	})

	t.Run("error cases", func(t *testing.T) {

		type testData struct {
			chain Blockchain
			txID  string
		}
		var testCases []testData
		for _, chain := range getTransactionBlockchains {
			testCases = append(testCases, testData{chain: chain, txID: testTxID(chain)})
		}

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": tx not found", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorTxNotFoundResponse{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": http req error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqErr{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": missing body contents", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqNoBodyErr{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": error with resp", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqWithRespErr{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": invalid json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadJSONResponse{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": invalid error json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadErrorJSONResponse{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})

			t.Run("chain "+testCase.chain.String()+": missing api key", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorMissingAPIKey{}))
				info, err := c.GetTransaction(context.Background(), testCase.chain, testCase.txID)
				require.Error(t, err)
				require.Nil(t, info)
			})
		}
	})
}

func ExampleClient_GetTransaction() {
	c := NewClient(WithHTTPClient(&validTxResponse{}))
	info, _ := c.GetTransaction(context.Background(), BSV, testTxID(BSV))
	fmt.Println("tx found: " + info.TxID)
	// Output:tx found: 17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd
}

func BenchmarkClient_GetTransaction(b *testing.B) {
	c := NewClient(WithHTTPClient(&validTxResponse{}))
	ctx := context.Background()
	tx := testTxID(BSV)
	for i := 0; i < b.N; i++ {
		_, _ = c.GetTransaction(ctx, BSV, tx)
	}
}

func TestClient_SendTransaction(t *testing.T) {
	t.Parallel()

	t.Run("valid cases", func(t *testing.T) {

		var tests = []struct {
			chain        Blockchain
			txHex        string
			expectedTxID string
		}{
			{BCH, testTxHex(BCH), testTxHexID(BCH)},
			{BSV, testTxHex(BSV), testTxHexID(BSV)},
			{BTC, testTxHex(BTC), testTxHexID(BTC)},
			{BTCTestnet, testTxHex(BTCTestnet), testTxHexID(BTCTestnet)},
			{BTG, testTxHex(BTG), testTxHexID(BTG)},
			{DASH, testTxHex(DASH), testTxHexID(DASH)},
			{DOGE, testTxHex(DOGE), testTxHexID(DOGE)},
			{LTC, testTxHex(LTC), testTxHexID(LTC)},
		}

		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()

		for _, testCase := range tests {
			t.Run("chain "+testCase.chain.String()+": SendTransaction("+testCase.txHex+")", func(t *testing.T) {
				results, err := c.SendTransaction(ctx, testCase.chain, testCase.txHex)
				require.NoError(t, err)
				require.NotNil(t, results)

				assert.Equal(t, testCase.expectedTxID, results.Result)
			})
		}
	})

	t.Run("missing or invalid tx hex", func(t *testing.T) {
		type testData struct {
			chain Blockchain
			txHex string
			err   error
		}
		var testCases []testData
		for _, chain := range sendTransactionBlockchains {
			testCases = append(testCases, testData{chain: chain, txHex: "", err: ErrInvalidTxHex})
			testCases = append(testCases, testData{chain: chain, txHex: "12345", err: ErrInvalidTxHex})
			testCases = append(testCases, testData{chain: chain, txHex: "invalid-tx-hex", err: ErrInvalidTxHex})
			testCases = append(testCases, testData{chain: chain, txHex: randomHexString(2001), err: ErrTxHexTooLarge})
		}

		c := NewClient(WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": invalid-tx", func(t *testing.T) {
				results, err := c.SendTransaction(ctx, testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
				assert.ErrorIs(t, err, testCase.err)
			})
		}
	})

	t.Run("unsupported chain", func(t *testing.T) {
		c := NewClient(WithHTTPClient(&validTxResponse{}))
		ctx := context.Background()
		sendTransactionBlockchains = []Blockchain{BSV}
		results, err := c.SendTransaction(ctx, BTC, testTxHex(BTC))
		require.Error(t, err)
		require.Nil(t, results)
		assert.ErrorIs(t, err, ErrUnsupportedBlockchain)
	})

	t.Run("error cases", func(t *testing.T) {

		type testData struct {
			chain Blockchain
			txHex string
		}
		var testCases []testData
		for _, chain := range sendTransactionBlockchains {
			testCases = append(testCases, testData{chain: chain, txHex: testTxHex(chain)})
		}

		for _, testCase := range testCases {
			t.Run("chain "+testCase.chain.String()+": broadcast error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorSendTxErrorResponse{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": http req error", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqErr{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": missing body contents", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqNoBodyErr{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": error with resp", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorDoReqWithRespErr{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": invalid json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadJSONResponse{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": invalid error json response", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorBadErrorJSONResponse{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})

			t.Run("chain "+testCase.chain.String()+": missing api key", func(t *testing.T) {
				c := NewClient(WithHTTPClient(&errorMissingAPIKey{}))
				results, err := c.SendTransaction(context.Background(), testCase.chain, testCase.txHex)
				require.Error(t, err)
				require.Nil(t, results)
			})
		}
	})
}

func ExampleClient_SendTransaction() {
	c := NewClient(WithHTTPClient(&validTxResponse{}))
	results, _ := c.SendTransaction(context.Background(), BSV, testTxHex(BSV))
	fmt.Println("broadcast success: " + results.Result)
	// Output:broadcast success: 15e78db3a6247ca320de2202240f6a4877ea3af338e23bf5ff3e5cbff3763bf6
}

func BenchmarkClient_SendTransaction(b *testing.B) {
	c := NewClient(WithHTTPClient(&validTxResponse{}))
	ctx := context.Background()
	tx := testTxHex(BSV)
	for i := 0; i < b.N; i++ {
		_, _ = c.SendTransaction(ctx, BSV, tx)
	}
}
