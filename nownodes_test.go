package nownodes

const (
	testBCHTxID     = "fb91f4b1969c7dcff3a1199b26ba9b33af658cec98641a32740f06f0e09b0efe"
	testBitcoinTxID = "17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd"
	testBTCTxID     = "050ead9fbd6360771541c734fa2324e5caa13f62a9598a39ca730aeb19e8c89a"
	testBTGTxID     = "934989d8e6e1fe9bc3d7508479df85e4757e6f0ceee613e7f8736e4e6b344a4a"
	testDASHTxID    = "7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e"
	testDOGETxID    = "6b22cc41b1206b6f39568bca5ca9e32ca3e0f6f4e0a68e2b126913e7d6620543"
	testKey         = "test-key-1234567"
	testLTCTxID     = "dfca839b7686a458e94001e53df4bd3bbe967d7ba5622f67b38cd6e2650bb37a"
	testUserAgent   = "test-user-agent"
)

func testTxID(chain Blockchain) string {
	switch chain {
	case BCH:
		return testBCHTxID
	case BTC:
		return testBTCTxID
	case BTG:
		return testBTGTxID
	case DASH:
		return testDASHTxID
	case DOGE:
		return testDOGETxID
	case LTC:
		return testLTCTxID
	default:
		return testBitcoinTxID
	}
}
