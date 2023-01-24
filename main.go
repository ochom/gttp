package gohttp

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"time"
)

// Request is the format required to make request
type Request struct {
	url     string
	headers map[string]string
	body    *bytes.Buffer
	timeout time.Duration
}

// New creates a new request with no configuration
func New() *Request {
	return &Request{}
}

// SetBody sets the body
func (r *Request) SetBody(body []byte) {
	r.body = bytes.NewBuffer(body)
}

// SetHeaders sets the headers
func (r *Request) SetHeaders(headers map[string]string) {
	r.headers = headers
}

// SetURL sets the url
func (r *Request) SetURL(url string) {
	r.url = url
}

// SetTimeout sets the timeout
func (r *Request) SetTimeout(seconds int) {
	r.timeout = time.Duration(seconds) * time.Second
}

// NewRequest creates a new request
func NewRequest(url string, headers map[string]string, body []byte) *Request {
	timeout := 10 * time.Second
	return &Request{url, headers, bytes.NewBuffer(body), timeout}
}

// NewRequestWithTimeout creates a new request with timeout in seconds
func NewRequestWithTimeout(url string, headers map[string]string, body []byte, seconds int) *Request {
	to := time.Duration(seconds) * time.Second
	return &Request{url, headers, bytes.NewBuffer(body), to}
}

// Send sends the request
func (r *Request) Send(method string) (body []byte, status int, err error) {
	timeout := 10 * time.Second
	if r.timeout == 0 {
		r.timeout = timeout
	}
	client := http.Client{
		Timeout: r.timeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: r.timeout,
			}).Dial,
			TLSHandshakeTimeout: r.timeout,
		},
	}

	req, err := http.NewRequest(method, r.url, r.body)
	if err != nil {
		return
	}

	for k, v := range r.headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	status = res.StatusCode

	return
}

// Post ...
func (r *Request) Post() (body []byte, status int, err error) {
	body, status, err = r.Send(http.MethodPost)
	return
}

// Get ...
func (r *Request) Get() (body []byte, status int, err error) {
	body, status, err = r.Send(http.MethodGet)
	return
}

// Put ...
func (r *Request) Put() (body []byte, status int, err error) {
	body, status, err = r.Send(http.MethodPut)
	return
}

// Patch ...
func (r *Request) Patch() (body []byte, status int, err error) {
	body, status, err = r.Send(http.MethodPatch)
	return
}
