package nownodes

import "errors"

// ErrInvalidTxID is when the tx id is missing or invalid
var ErrInvalidTxID = errors.New("missing or invalid tx id")

// ErrUnsupportedBlockchain is when the given blockchain is not supported by the method
var ErrUnsupportedBlockchain = errors.New("unsupported blockchain for this method")
