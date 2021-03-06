package nownodes

import "errors"

// ErrInvalidTxID is when the tx id is missing or invalid
var ErrInvalidTxID = errors.New("missing or invalid tx id")

// ErrInvalidTxHex is when the tx hex is missing or invalid
var ErrInvalidTxHex = errors.New("missing or invalid tx hex")

// ErrInvalidAddress is when the address is missing or invalid
var ErrInvalidAddress = errors.New("missing or invalid address")

// ErrUnsupportedBlockchain is when the given blockchain is not supported by the method
var ErrUnsupportedBlockchain = errors.New("unsupported blockchain for this method")
