package nownodes

const (
	// version is the current version
	version = "v0.0.3"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-nownodes: " + version

	// apiVersion is the current NOWNodes API version
	apiVersion  = "v2"
	nowNodesURL = "nownodes.io"

	// Appends all requests with this protocol
	httpProtocol = "https://"

	// API header key for NOWNodes API
	apiHeaderKey = "api-key"

	// Coin specific values
	bitcoinCashPrefix = "bitcoincash:"

	// Bitcoin transaction length
	bitcoinCashMaxAddressLength = 42
	bitcoinMaxAddressLength     = 35
	bitcoinMinAddressLength     = 26
	bitcoinTransactionLength    = 64
	liteCoinMaxAddressLength    = 43
	maxTxHexLengthOnSend        = 2000

	// Blockchains
	blockchainBCH        = "bch"
	blockchainBSV        = "bsv"
	blockchainBTC        = "btc"
	blockchainBTCTestnet = "btc-testnet"
	blockchainBTG        = "btg"
	blockchainDASH       = "dash"
	blockchainDOGE       = "doge"
	blockchainLTC        = "ltc"

	// Routes
	routeGetAddress = "/address/"
	routeGetTx      = "/tx/"
	routeSendTx     = "/sendtx/"
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
	getTransactionBlockchains = allBlockchains

	// Supported blockchains for the method GetAddress()
	getAddressBlockchains = getTransactionBlockchains

	// Supported blockchains for the method SendTransaction()
	sendTransactionBlockchains = getTransactionBlockchains
)
