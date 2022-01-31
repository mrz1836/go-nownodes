package nownodes

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

const (
	// Testing variables
	testKey       = "test-key-1234567"      // Test API Key
	testUniqueID  = "test-custom-unique-id" // Test id for node requests
	testUserAgent = "test-user-agent"       // Test user agent

	// Test transactions
	testBCHTxID     = "fb91f4b1969c7dcff3a1199b26ba9b33af658cec98641a32740f06f0e09b0efe"   // https://blockchair.com/bitcoin-cash/transaction/<txid>
	testBitcoinTxID = "17961a51337369bf64e45e8410a7ce4cfb0c88b5d883d9e8a939dfdd0f7591fd"   // https://blockchair.com/bitcoin-sv/transaction/<txid>
	testBTCTxID     = "050ead9fbd6360771541c734fa2324e5caa13f62a9598a39ca730aeb19e8c89a"   // https://blockchair.com/bitcoin/transaction/<txid>
	testBTGTxID     = "934989d8e6e1fe9bc3d7508479df85e4757e6f0ceee613e7f8736e4e6b344a4a"   // https://explorer.bitcoingold.org/insight/tx/<txid>
	testDASHTxID    = "7c4738c76ba318e74af3e6f85f94ac15b44b4189bf76a9523aed71366a75445e"   // https://blockchair.com/dash/transaction/<txid>
	testDOGETxID    = "6b22cc41b1206b6f39568bca5ca9e32ca3e0f6f4e0a68e2b126913e7d6620543"   // https://blockchair.com/dogecoin/transaction/<txid>
	testLTCTxID     = "dfca839b7686a458e94001e53df4bd3bbe967d7ba5622f67b38cd6e2650bb37a"   // https://blockchair.com/litecoin/transaction/<txid>
	testETHTxID     = "0x193f9293a30bf668e3edd2290d49534770e6118faa201fe97da498cd8a995765" // https://etherscan.io/tx/<txid>

	// Test addresses
	testBitcoinAddress = "1GenocdBC1NSHLMbk61fqJXqTdXjevCxCL"                     // https://blockchair.com/bitcoin-sv/address/<address>
	testBCHAddress     = "bitcoincash:qzgztrce3qtc272dfffzc0lz3e02ykvunyzaud5kdn" // https://blockchair.com/bitcoin-cash/address/<address>
	testBTCAddress     = "17eKje3fzPs633GKsotMFkLKRzv1HPSRTz"                     // https://blockchair.com/bitcoin/address/<address>
	testBTGAddress     = "ATTav2PtmotBZwxgWjrZCgaZpE89kcJ29B"                     // https://explorer.bitcoingold.org/insight/address/<address>
	testDASHAddress    = "Xe7tPVUvDpt52h2KykMpj4hmh8VsAaPCgt"                     // https://blockchair.com/dash/address/<address>
	testDOGEAddress    = "ACSbgj91BsjdpuG6pBkG9LXtCTFaH4mn5a"                     // https://blockchair.com/dogecoin/address/<address>
	testLTCAddress     = "ltc1q684prk03qaz2uqenylktaq5z9qcg3dz7fgg0ps"            // https://blockchair.com/litecoin/address/<address>
	testETHAddress     = "0x7dbf304559293bdccac7cb18cb69375719b5e1cc"             // https://etherscan.io/address/<address>

	// Test transaction hex & id for send transaction
	testBitcoinTxHex   = "01000000017f04d780417bff05f6ca3b210c86308e848b9295f449a89d5091861573cebfc1020000006b483045022100f160e411a9a2c3b9fd975c0f8807186f6dfbfb9bd29a4ce272689dc14ff0d1ff0220092176eca8e284ebbabd534b3812942e14c9d655c5fcaa08740746a770606088412103791d1cf1ec22e86006b42a5b6fba7312a19d1578ae919aa2f4dd3d8d0c04ced8ffffffff020000000000000000b4006a0372756e0105036679784ca67b22696e223a312c22726566223a5b5d2c226f7574223a5b5d2c2264656c223a5b2237313237323035363434633631366361626365323161366432383033663038356538373038653966356635326135663838326432303835663333303966373961225d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c2264657374726f79222c5b5d5d7d5d7dc2010000000000001976a9147e7d79a417a21c125c43552446cb4aadb5d41c1188ac00000000"
	testBitcoinTxHexID = "15e78db3a6247ca320de2202240f6a4877ea3af338e23bf5ff3e5cbff3763bf6"
	testBTCTxHex       = "020000000001013fee71f3b62b871b8f50e52e2c408ae9282e6896030ec15d3e2e4101248a14100100000000fdffffff02cc8b000000000000160014ef0c54cd24cd6036662dab76123d133b99e5842395a300000000000017a914200f6d0d50c82713ac1543044d695a04180ec5978702473044022079cb1b845c509cac67b952170c7b2fe061729e6402767a2706398b4f8596a9160220604f3b3d64c68cd43c428da9ed480bb2d152cf7e451e26a6d9b04f818b672fb1012103eee9326b3c204620124ab38415a5aa152eef2db2cf8a6457a72b803a5dae543a3e010b00"
	testBTCTxHexID     = "4e475f486c5aad520d113448f70980f40e9b4976e4664bb3a3ed7d14c00ba639"
	testBTGTxHex       = "0100000001350aeec47611bc79232575f4cf22fc037005383782d2e03279d09096863da1f4010000006b483045022100c5a4c7bcaef385e93ac7ddfcd8eba132d2f0166921f98a59c1e942a93e8056e002200bd2b8b830c689f7d37d1a7b9bda300d55d59d9be32e0c2bbd28babe005cff1341210245766ba2b274073a604fe17fc44ffb25e5927142fc58c48ac3f59ca96ca330fdffffffff0123020000000000001976a9147b7385c6632ed95b06afc1e3240780ddbe4893d888ac00000000"
	testBTGTxHexID     = "683e11d4db8a776e293dc3bfe446edf66cf3b145a6ec13e1f5f1af6bb5855364"
	testETHTxHex       = "0x02f8b10145843b9aca00852d08b84a94830350ad94a1c13e02a8b3f833d7b47fa57ba6484f656ee06780b84410abbfae00000000000000000000000002f30927eb29f3f66031517bb3de3948f32ee01900000000000000000000000000000000000000000000000000000000000001f4c001a04b5bd382c480a1b4ebf5a203ea482c3fd17bed03e9b810219dc79159aefb01c3a06b1a3dc875d99282b62b043b14d1f2a9f8867329b4e66719c622dbc315c5f0ea\n"
	testETHTxHexID     = "0x193f9293a30bf668e3edd2290d49534770e6118faa201fe97da498cd8a995765"
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
	case ETH:
		return testETHTxID
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
	case ETH:
		return testETHAddress
	default:
		return testBitcoinAddress
	}
}

func testTxHex(chain Blockchain) string {
	switch chain {
	case BCH:
		// note: add a real BCH tx
		return testBitcoinTxHex
	case BTC, BTCTestnet:
		return testBTCTxHex
	case BTG:
		return testBTGTxHex
	case DASH:
		// note: add a real DASH tx
		return testBitcoinTxHex
	case DOGE:
		// note: add a real DOGE tx
		return testBitcoinTxHex
	case LTC:
		// note: add a real LTC tx
		return testBitcoinTxHex
	case BSV:
		return testBitcoinTxHex
	case ETH:
		return testETHTxHex
	default:
		return testBitcoinTxHex
	}
}

func testTxHexID(chain Blockchain) string {
	switch chain {
	case BCH:
		return testBitcoinTxHexID
	case BTC, BTCTestnet:
		return testBTCTxHexID
	case BTG:
		return testBTGTxHexID
	case DASH:
		return testBitcoinTxHexID
	case DOGE:
		return testBitcoinTxHexID
	case LTC:
		return testBitcoinTxHexID
	case BSV:
		return testBitcoinTxHexID
	case ETH:
		return testETHTxHexID
	default:
		return testBitcoinTxHexID
	}
}

func randomHexString(length int) string {
	buff := make([]byte, int(math.Ceil(float64(length)/1.33333333333)))
	_, _ = rand.Read(buff)
	return hex.EncodeToString(buff)
}
