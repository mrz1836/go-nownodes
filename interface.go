package nownodes

import "context"

// AddressService is the address related requests
type AddressService interface {
	GetAddress(ctx context.Context, chain Blockchain, address string) (*AddressInfo, error)
}

// TransactionService is the transaction related requests
type TransactionService interface {
	GetTransaction(ctx context.Context, chain Blockchain, txID string) (*TransactionInfo, error)
}

// ClientInterface is the client interface
type ClientInterface interface {
	AddressService
	TransactionService
	HTTPClient() HTTPInterface
	UserAgent() string
}
