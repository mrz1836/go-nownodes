package nownodes

const (
	// version is the current version
	version = "v0.0.1"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-nownodes: " + version

	// apiVersion is the current NOWNodes API version
	apiVersion  = "v2"
	nowNodesURL = "nownodes.io"

	// Appends all requests with this protocol
	httpProtocol = "https://"

	// API header key for NOWNodes API
	apiHeaderKey = "api-key"

	// Bitcoin transaction length
	bitcoinTransactionLength = 64

	// Blockchains
	blockchainBCH        = "bch"
	blockchainBSV        = "bsv"
	blockchainBTC        = "btc"
	blockchainBTCTestnet = "btc-testnet"
	blockchainBTG        = "btg"
	blockchainDASH       = "dash"
	blockchainDOGE       = "doge"
	blockchainLTC        = "ltc"
)

var (

	// All blockchains (used in tests and listing methods)
	allBlockchains = []Blockchain{
		BCH,
		BSV,
		BTC,
		BTCTestnet,
		BTG,
		DASH,
		DOGE,
		LTC,
	}

	// Supported blockchains for the method GetTransaction()
	getTransactionBlockchains = []Blockchain{
		BCH,
		BSV,
		BTC,
		BTCTestnet,
		BTG,
		DASH,
		DOGE,
		LTC,
	}
)
