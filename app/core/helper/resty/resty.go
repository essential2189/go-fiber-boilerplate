package resty

import (
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	Client *resty.Client
}

type Response struct {
	Body       []byte
	StatusCode int
	Status     string
}

func NewRestyClient() *HttpClient {
	r := resty.New()

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Panicf("defaultRoundTripper not an *http.Transport")
	}

	defaultTransport := *defaultTransportPointer
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	r.SetTimeout(defaultTimeOut)
	r.SetRetryWaitTime(defaultRetryBackOff)

	return &HttpClient{
		Client: r,
	}
}
