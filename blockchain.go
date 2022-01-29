package nownodes

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

// ValidateTx will do basic validations on the tx id string
func (n Blockchain) ValidateTx(txID string) bool {
	switch n {
	case BCH:
		return len(txID) == bitcoinTransactionLength
	case BSV:
		return len(txID) == bitcoinTransactionLength
	case BTC:
		return len(txID) == bitcoinTransactionLength
	case BTCTestnet:
		return len(txID) == bitcoinTransactionLength
	case BTG:
		return len(txID) == bitcoinTransactionLength
	case DASH:
		return len(txID) == bitcoinTransactionLength
	case DOGE:
		return len(txID) == bitcoinTransactionLength
	case LTC:
		return len(txID) == bitcoinTransactionLength
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
