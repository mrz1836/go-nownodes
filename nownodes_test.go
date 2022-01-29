package nownodes

const (
	testKey       = "test-key-1234567" // Test API Key
	testUserAgent = "test-user-agent"  // Test user agent

	// Test transactions
	testBCHTxID     = "fb91f4b1969c7dcff3a1199b26ba9b33af658cec98641a32740f06f0e09b0efe" // https://blockchair.com/bitcoin-cash/transaction/<txid>
	testBitcoinTxID = "17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd" // https://blockchair.com/bitcoin-sv/transaction/<txid>
	testBTCTxID     = "050ead9fbd6360771541c734fa2324e5caa13f62a9598a39ca730aeb19e8c89a" // https://blockchair.com/bitcoin/transaction/<txid>
	testBTGTxID     = "934989d8e6e1fe9bc3d7508479df85e4757e6f0ceee613e7f8736e4e6b344a4a" // https://explorer.bitcoingold.org/insight/tx/<txid>
	testDASHTxID    = "7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e" // https://blockchair.com/dash/transaction/<txid>
	testDOGETxID    = "6b22cc41b1206b6f39568bca5ca9e32ca3e0f6f4e0a68e2b126913e7d6620543" // https://blockchair.com/dogecoin/transaction/<txid>
	testLTCTxID     = "dfca839b7686a458e94001e53df4bd3bbe967d7ba5622f67b38cd6e2650bb37a" // https://blockchair.com/litecoin/transaction/<txid>

	// Test addresses
	testBitcoinAddress = "1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL"                     // https://blockchair.com/bitcoin-sv/address/<address>
	testBCHAddress     = "bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn" // https://blockchair.com/bitcoin-cash/address/<address>
	testBTCAddress     = "17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"                     // https://blockchair.com/bitcoin/address/<address>
	testBTGAddress     = "ATTav2PtmotBZwxgWjrZCgaZpE89kcJ29B"                     // https://explorer.bitcoingold.org/insight/address/<address>
	testDASHAddress    = "Xe7tPVUvDpt52h2KykMpj4hmh8VsAaPCgt"                     // https://blockchair.com/dash/address/<address>
	testDOGEAddress    = "ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a"                     // https://blockchair.com/dogecoin/address/<address>
	testLTCAddress     = "ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps"            // https://blockchair.com/litecoin/address/<address>
)

func testTxID(chain Blockchain) string {
	switch chain {
	case BCH:
		return testBCHTxID
	case BTC, BTCTestnet:
		return testBTCTxID
	case BTG:
		return testBTGTxID
	case DASH:
		return testDASHTxID
	case DOGE:
		return testDOGETxID
	case LTC:
		return testLTCTxID
	case BSV:
		return testBitcoinTxID
	default:
		return testBitcoinTxID
	}
}

func testAddress(chain Blockchain) string {
	switch chain {
	case BCH:
		return testBCHAddress
	case BTC, BTCTestnet:
		return testBTCAddress
	case BTG:
		return testBTGAddress
	case DASH:
		return testDASHAddress
	case DOGE:
		return testDOGEAddress
	case LTC:
		return testLTCAddress
	case BSV:
		return testBitcoinAddress
	default:
		return testBitcoinAddress
	}
}
