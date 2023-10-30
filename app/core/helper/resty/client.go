package resty

import (
	"time"
)

type RequestInfo struct {
	Uri          string
	Method       string
	Headers      map[string]string
	Query        map[string]string
	Body         interface{}
	Timeout      time.Duration
	RetryCount   int
	RetryBackOff time.Duration
	IsSkipSSL    bool
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
}
