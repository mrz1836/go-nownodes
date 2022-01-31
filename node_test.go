package nownodes

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// validNodeResponse will return a valid response for all supported blockchains
type validNodeResponse struct{}

func (v *validNodeResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	decoder := json.NewDecoder(req.Body)
	var data nodePayload
	err := decoder.Decode(&data)
	if err != nil {
		return resp, err
	}

	// Valid response (get tx)
	for _, chain := range sendRawTransactionBlockchains {
		if strings.Contains(req.Host, chain.NodeAPIURL()) && data.Method == nodeMethodSendRawTx {
			resp.StatusCode = http.StatusOK
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"result": "` + testTxHexID(chain) + `","error": null,"id": "` + data.ID + `"}`)))
			return resp, nil
		}
	}

	// Valid response (get mempool entry)
	for _, chain := range getMempoolEntryBlockchains {
		if strings.Contains(req.Host, chain.NodeAPIURL()) && data.Method == nodeMethodGetMempoolEntry {
			resp.StatusCode = http.StatusOK
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"result": {"size": 381,"fee": 9.6e-7,"modifiedfee": 9.6e-7,"time": 1643661192,"height": 724704,"depends": []},"error": null,"id": "` + data.ID + `"}`)))
			return resp, nil
		}
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}

// errorNodeErrorResponse will return an error for the "send raw tx" response
type errorNodeErrorResponse struct{}

func (v *errorNodeErrorResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, errors.New("missing request")
	}

	decoder := json.NewDecoder(req.Body)
	var data nodePayload
	err := decoder.Decode(&data)
	if err != nil {
		return resp, err
	}

	// Error response (send tx)
	for _, chain := range sendRawTransactionBlockchains {
		if strings.Contains(req.Host, chain.NodeAPIURL()) && data.Method == nodeMethodSendRawTx {
			resp.StatusCode = http.StatusInternalServerError
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"result": null,"error": {"code": -27,"message": "Transaction already in the mempool"},"id": "` + data.ID + `"}`)))
			return resp, nil
		}
	}

	// Error response (get mempool entry)
	for _, chain := range getMempoolEntryBlockchains {
		if strings.Contains(req.Host, chain.NodeAPIURL()) && data.Method == nodeMethodGetMempoolEntry {
			resp.StatusCode = http.StatusInternalServerError
			resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"result": null,"error": {"code": -5,"message": "Transaction not in mempool"},"id": "` + data.ID + `"}`)))
			return resp, nil
		}
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"no-route-found"}`)))
	return resp, errors.New("request not found")
}
