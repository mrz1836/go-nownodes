package nownodes

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// errorDoReqErr will return an error for the HTTP request
type errorDoReqErr struct{}

func (v *errorDoReqErr) Do(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("http error or Do() error")
}

// errorDoReqNoBodyErr will return an error for the HTTP request
type errorDoReqNoBodyErr struct{}

func (v *errorDoReqNoBodyErr) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest
	return resp, nil
}

// errorDoReqWithRespErr will return an error for the HTTP request
type errorDoReqWithRespErr struct{}

func (v *errorDoReqWithRespErr) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest
	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(``)))
	return resp, errors.New("http error or Do() error")
}

// errorBadJSONResponse will return an error for bad json
type errorBadJSONResponse struct{}

func (v *errorBadJSONResponse) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusOK
	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"invalid":json}`)))
	return resp, nil
}

// errorBadErrorJSONResponse will return an error for bad json
type errorBadErrorJSONResponse struct{}

func (v *errorBadErrorJSONResponse) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest
	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"invalid":json}`)))
	return resp, nil
}

// errorMissingAPIKey will return an error for bad json
type errorMissingAPIKey struct{}

func (v *errorMissingAPIKey) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusOK
	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"message": "Missing api-key header"}`)))
	// {"message": "Unknown API_key"}
	return resp, nil
}
