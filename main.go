package gohttp

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

//RequestPayload is the format required to make request
type RequestPayload struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    *bytes.Buffer
}

//Service blueprint to the available functions
type Service interface {
	// Post makes a post request returns status code, response body and any errors
	Post(ctx context.Context, url string, headers map[string]string, payload []byte) (int, []byte, error)

	// Get makes a get request returns status code, response body and any errors
	Get(ctx context.Context, url string, headers map[string]string) (int, []byte, error)

	//PostHere post to posthere.io
	PostHere(ctx context.Context, data []byte)
}

//impl the controller struct for external access
type impl struct {
	client http.Client
}

//New create a new instance of HTTPService
func New(timeout time.Duration) Service {
	client := http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: timeout,
			}).Dial,
			TLSHandshakeTimeout: timeout,
		},
	}
	return &impl{
		client: client,
	}
}

//Post ...
func (s *impl) Post(ctx context.Context, url string, headers map[string]string, payload []byte) (status int, body []byte, err error) {
	status = 0
	body = []byte{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	status = res.StatusCode

	return
}

// Get ...
func (s *impl) Get(ctx context.Context, url string, headers map[string]string) (status int, body []byte, err error) {
	status = 0
	body = []byte{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	status = res.StatusCode

	return
}

//PostHere ...
func (s *impl) PostHere(ctx context.Context, data []byte) {
	url := "https://posthere.io/f8c4-4160-b821"

	_, _, err := s.Post(ctx, url, nil, data)
	if err != nil {
		log.Println(err.Error())
	}
}
