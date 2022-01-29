package nownodes

import "time"

// ClientOps allow functional options to be supplied that overwrite default client options.
type ClientOps func(c *ClientOptions)

// DefaultHTTPOptions will return the default HTTP option values
func DefaultHTTPOptions() (clientOptions *HTTPOptions) {
	return &HTTPOptions{
		BackOffExponentFactor:          2.0,
		BackOffInitialTimeout:          2 * time.Millisecond,
		BackOffMaximumJitterInterval:   2 * time.Millisecond,
		BackOffMaxTimeout:              10 * time.Millisecond,
		DialerKeepAlive:                20 * time.Second,
		DialerTimeout:                  5 * time.Second,
		RequestRetryCount:              2,
		RequestTimeout:                 30 * time.Second,
		TransportExpectContinueTimeout: 3 * time.Second,
		TransportIdleTimeout:           20 * time.Second,
		TransportMaxIdleConnections:    10,
		TransportTLSHandshakeTimeout:   5 * time.Second,
	}
}

// WithAPIKey will store the API key on the client for all future requests
func WithAPIKey(apiKey string) ClientOps {
	return func(c *ClientOptions) {
		if len(apiKey) > 0 {
			c.apiKey = apiKey
		}
	}
}

// WithHTTPClient will overwrite the default client with a custom client
func WithHTTPClient(client HTTPInterface) ClientOps {
	return func(c *ClientOptions) {
		if client != nil {
			c.httpClient = client
		}
	}
}

// WithHTTPOptions will overwrite the default HTTP client options
func WithHTTPOptions(opts *HTTPOptions) ClientOps {
	return func(c *ClientOptions) {
		if opts != nil {
			c.httpOptions = opts
		}
	}
}

// WithUserAgent will overwrite the default useragent
func WithUserAgent(userAgent string) ClientOps {
	return func(c *ClientOptions) {
		if len(userAgent) > 0 {
			c.userAgent = userAgent
		}
	}
}
