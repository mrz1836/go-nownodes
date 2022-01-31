package nownodes

import "context"

// AddressService is the address related requests
type AddressService interface {
	GetAddress(ctx context.Context, chain Blockchain, address string) (*AddressInfo, error)
}

// MempoolService is the mempool related requests
type MempoolService interface {
	GetMempoolEntry(ctx context.Context, chain Blockchain, txID, id string) (*MempoolEntryResult, error)
}

// TransactionService is the transaction related requests
type TransactionService interface {
	GetTransaction(ctx context.Context, chain Blockchain, txID string) (*TransactionInfo, error)
	SendTransaction(ctx context.Context, chain Blockchain, txHex string) (*BroadcastResult, error)
	SendRawTransaction(ctx context.Context, chain Blockchain, txHex, id string) (*BroadcastResult, error)
}

// ClientInterface is the client interface
type ClientInterface interface {
	AddressService
	MempoolService
	TransactionService
	HTTPClient() HTTPInterface
	UserAgent() string
}
