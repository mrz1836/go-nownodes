package nownodes

import "context"

// TransactionService is the transaction related requests
type TransactionService interface {
	GetTransaction(ctx context.Context, chain Blockchain, txID string) (*TransactionInfo, error)
}

// ClientInterface is the client interface
type ClientInterface interface {
	TransactionService
	HTTPClient() HTTPInterface
	UserAgent() string
}
