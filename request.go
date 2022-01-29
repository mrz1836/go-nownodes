package nownodes

import (
	"context"
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
	// Data   []byte `json:"data"`
	APIKey string `json:"api_key"`
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
	/*
		todo: enable once there are requests that require these methods
		if payload.Method == http.MethodPost || payload.Method == http.MethodPut {
			bodyReader = bytes.NewBuffer(payload.Data)
			response.PostData = string(payload.Data)
		}
	*/

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
	/*
		note: see above
		if payload.Method == http.MethodPost || payload.Method == http.MethodPut {
			request.Header.Set("Content-Type", "application/json")
		}
	*/

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
	return
}
