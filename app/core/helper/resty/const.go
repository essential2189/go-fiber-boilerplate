package resty

import "time"

const (
	defaultRetryBackOff = 100 * time.Millisecond
	defaultTimeOut      = 10 * time.Second
	MethodGET           = "GET"
	MethodPOST          = "POST"
	MethodDELETE        = "DELETE"
	MethodPUT           = "PUT"
)
