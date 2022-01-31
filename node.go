package nownodes

// NodeError is an internal error from the NodeAPI
type NodeError struct {
	Error *nodeAPIError `json:"error,omitempty"` // The error message from NodeAPI requests
}

// nodeError is an internal error from the NodeAPI
type nodeAPIError struct {
	Code    int64  `json:"code"`    // IE: -26
	Message string `json:"message"` // IE: 257: txn-already-known
}
