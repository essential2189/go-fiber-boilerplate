package resty

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

type Request interface {
	Request(r RequestInfo) (*HttpResponse, error)
}

func (c *HttpClient) Request(r RequestInfo) (*HttpResponse, error) {
	resp, err := request(c, r.Uri, r.Method, r.Headers, r.Query, r.Body, r.Timeout, r.RetryCount, r.RetryBackOff, r.IsSkipSSL)
	if err != nil {
		return nil, err
	}

	body := &HttpResponse{StatusCode: resp.StatusCode, Body: resp.Body}
	return body, nil
}

func request(h *HttpClient, uri, method string, header, query map[string]string, body interface{}, timeout time.Duration, retryCount int, retryBackOff time.Duration, isSkipSSL bool) (*Response, error) {

	if timeout > 0 {
		h.Client.SetTimeout(timeout)
	}

	if retryCount > 0 {
		h.Client.SetRetryCount(retryCount)
	}

	if retryBackOff > 0 {
		h.Client.SetRetryWaitTime(retryBackOff)
	}

	if isSkipSSL {
		h.Client.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	r := h.Client.R()

	r.SetQueryParams(query)
	r.SetBody(body)
	r.SetHeaders(header)

	var resp *resty.Response
	var err error

	switch method {
	case MethodGET:
		resp, err = r.Get(uri)
	case MethodPOST:
		resp, err = r.Post(uri)
	case MethodPUT:
		resp, err = r.Put(uri)
	case MethodDELETE:
		resp, err = r.Delete(uri)
	default:
		resp, err = r.Execute(method, uri)
	}
	if err != nil {
		return nil, err
	}

	return &Response{
		Status:     resp.Status(),
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
	}, nil
}
