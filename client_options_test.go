package nownodes

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultHTTPOptions(t *testing.T) {
	t.Parallel()

	options := DefaultHTTPOptions()

	assert.Equal(t, 10, options.TransportMaxIdleConnections)
	assert.Equal(t, 2*time.Millisecond, options.BackOffInitialTimeout)
	assert.Equal(t, 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	assert.Equal(t, 2, options.RequestRetryCount)
	assert.Equal(t, 2.0, options.BackOffExponentFactor)
	assert.Equal(t, 20*time.Second, options.DialerKeepAlive)
	assert.Equal(t, 20*time.Second, options.TransportIdleTimeout)
	assert.Equal(t, 3*time.Second, options.TransportExpectContinueTimeout)
	assert.Equal(t, 30*time.Second, options.RequestTimeout)
	assert.Equal(t, 5*time.Second, options.DialerTimeout)
	assert.Equal(t, 5*time.Second, options.TransportTLSHandshakeTimeout)
}

func TestWithAPIKey(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithAPIKey("")
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying empty", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithAPIKey("")
		opt(options)
		assert.Equal(t, "", options.apiKey)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithAPIKey(testKey)
		opt(options)
		assert.Equal(t, testKey, options.apiKey)
	})
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithHTTPClient(nil)
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying nil", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithHTTPClient(nil)
		opt(options)
		assert.Nil(t, options.httpOptions)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		customClient := &http.Client{}
		opt := WithHTTPClient(customClient)
		opt(options)
		assert.Equal(t, customClient, options.httpClient)
	})
}

func TestWithHTTPOptions(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithHTTPOptions(nil)
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying nil", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithHTTPOptions(nil)
		opt(options)
		assert.Nil(t, options.httpOptions)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		customHTTPOpts := DefaultHTTPOptions()
		customHTTPOpts.RequestRetryCount = 3
		opt := WithHTTPOptions(customHTTPOpts)
		opt(options)
		assert.Equal(t, customHTTPOpts, options.httpOptions)
	})
}

func TestWithUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithUserAgent("")
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying empty", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithUserAgent("")
		opt(options)
		assert.Equal(t, "", options.userAgent)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithUserAgent(testUserAgent)
		opt(options)
		assert.Equal(t, testUserAgent, options.userAgent)
	})
}
