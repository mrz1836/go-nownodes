package nownodes

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// RequestResponse is the response from a request
type RequestResponse struct {
	BodyContents []byte `json:"body_contents"` // Raw body response
	Error        error  `json:"error"`         // If an error occurs
	Method       string `json:"method"`        // Method is the HTTP method used
	PostData     string `json:"post_data"`     // PostData is the post data submitted if POST/PUT request
	StatusCode   int    `json:"status_code"`   // StatusCode is the last code from the request
	URL          string `json:"url"`           // URL is used for the request
}

// httpPayload is used for a httpRequest
type httpPayload struct {
	APIKey string `json:"api_key"`
	Data   []byte `json:"data"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

// httpRequest is a generic request wrapper that can be used without constraints
func httpRequest(ctx context.Context, client *Client,
	payload *httpPayload) (response *RequestResponse) {

	// Set reader & response
	var bodyReader io.Reader
	response = new(RequestResponse)

	// Add post data if applicable
	if payload.Method == http.MethodPost {
		bodyReader = bytes.NewBuffer(payload.Data)
		response.PostData = string(payload.Data)
	}

	// Store for debugging purposes
	response.Method = payload.Method
	response.URL = payload.URL

	// Start the request
	var request *http.Request
	if request, response.Error = http.NewRequestWithContext(
		ctx, payload.Method, payload.URL, bodyReader,
	); response.Error != nil {
		return
	}

	// Change the header (user agent is in case they block default Go user agents)
	request.Header.Set("User-Agent", client.options.userAgent)

	// Set the content type on Method
	if payload.Method == http.MethodPost {
		request.Header.Set("Content-Type", "application/json")
	}

	// Set a token if supplied
	if len(payload.APIKey) > 0 {
		request.Header.Set(apiHeaderKey, payload.APIKey)
	}

	// Fire the http request
	var resp *http.Response
	if resp, response.Error = client.options.httpClient.Do(request); response.Error != nil {
		if resp != nil {
			response.StatusCode = resp.StatusCode
		}
		return
	}

	// Close the response body
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	// Set the status
	response.StatusCode = resp.StatusCode

	// Read the body
	if resp.Body != nil {
		response.BodyContents, response.Error = ioutil.ReadAll(resp.Body)
	}

	// Check status code
	if http.StatusOK == resp.StatusCode {

		// Detect error message (it's not an error, and returns a 200)
		if strings.Contains(string(response.BodyContents), `{"message":`) {
			errBody := struct {
				Error string `json:"message"`
			}{}
			if err := json.Unmarshal(
				response.BodyContents, &errBody,
			); err != nil {
				response.Error = fmt.Errorf("failed to unmarshal error response: %w", err)
				return
			}
			response.Error = errors.New(errBody.Error)
		}
		return
	}

	// There's no "body" present, so just echo status code.
	if response.BodyContents == nil {
		response.Error = fmt.Errorf(
			"status code: %d does not match %d",
			resp.StatusCode, http.StatusOK,
		)
		return
	}

	// Have a "body" so map to an error type and add to the error message.
	if payload.Method == http.MethodGet {
		errBody := struct {
			Error string `json:"error"`
		}{}
		if err := json.Unmarshal(
			response.BodyContents, &errBody,
		); err != nil {
			response.Error = fmt.Errorf("failed to unmarshal error response: %w", err)
			return
		}
		response.Error = errors.New(errBody.Error)
	} else {
		errBody := new(NodeError)
		if err := json.Unmarshal(
			response.BodyContents, &errBody,
		); err != nil {
			response.Error = fmt.Errorf("failed to unmarshal error response: %w", err)
			return
		}
		response.Error = fmt.Errorf(
			"code [%d] error [%s]", errBody.Error.Code, errBody.Error.Message,
		)
	}
	return
}

// blockBookRequest will make a BlockBook request and imbue the results into the given model
func blockBookRequest(ctx context.Context, client *Client, chains []Blockchain,
	chain Blockchain, endpoint string, model interface{}) error {

	resp, err := blockBookRequestInternal(
		ctx, client, chains, chain, endpoint,
	)
	if err != nil {
		return err
	}

	// Unmarshal the response
	return json.Unmarshal(
		resp.BodyContents, &model,
	)
}

// blockBookRequestInternal will make a BlockBook request and return the result
func blockBookRequestInternal(ctx context.Context, client *Client, chains []Blockchain,
	chain Blockchain, endpoint string) (*RequestResponse, error) {

	// Are we using a supported blockchain?
	if !isBlockchainSupported(chains, chain) {
		return nil, ErrUnsupportedBlockchain
	}

	// Fire the HTTP request
	resp := httpRequest(ctx, client, &httpPayload{
		APIKey: client.options.apiKey,
		Method: http.MethodGet,
		URL:    httpProtocol + chain.BlockBookURL() + "/api/" + apiVersion + endpoint,
	})
	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp, nil
}

/*
// blockBookRequestWithNoResponse will make a BlockBook request and only return an error if it fails
func blockBookRequestWithNoResponse(ctx context.Context, client *Client, chains []Blockchain,
	chain Blockchain, endpoint string) error {

	_, err := blockBookRequestInternal(
		ctx, client, chains, chain, endpoint,
	)
	if err != nil {
		return err
	}
	return nil
}
*/

// nodeRequest will make a NodeAPI request and return the result
func nodeRequest(ctx context.Context, client *Client, chains []Blockchain,
	chain Blockchain, payload []byte, model interface{}) error {

	// Are we using a supported blockchain?
	if !isBlockchainSupported(chains, chain) {
		return ErrUnsupportedBlockchain
	}

	// Fire the HTTP request
	resp := httpRequest(ctx, client, &httpPayload{
		APIKey: client.options.apiKey,
		Data:   payload,
		Method: http.MethodPost,
		URL:    httpProtocol + chain.NodeAPIURL(),
	})
	if resp.Error != nil {
		return resp.Error
	}

	// Unmarshal the response
	return json.Unmarshal(
		resp.BodyContents, &model,
	)
}

// nodePayload is the internal raw node payload
type nodePayload struct {
	APIKey  string   `json:"API_key"`
	ID      string   `json:"id"`
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

// createPayload will create the JSON payload for the NodeAPI requests
func createPayload(apiKey, method, id string, params []string) []byte {
	b, _ := json.Marshal(nodePayload{ // nolint: errchkjson // not going to produce an error
		APIKey:  apiKey,
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	})
	return b
}

// hashString will generate a hash of the given string
func hashString(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
