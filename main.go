package gohttp

import (
	"bytes"
	"context"
	"fmt"
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
	Body    []byte
}

//HTTPService blueprint to the available functions
type HTTPService interface {
	MakeRequest(ctx context.Context, payload RequestPayload) ([]byte, error)

	//post to posthere.io
	PostHere(ctx context.Context, headers map[string]string, data string)
}

//Service the controller struct for external access
type Service struct {
	Client http.Client
}

//NewHTTPService create a new instance of HTTPService
func NewHTTPService(timeout time.Duration) HTTPService {
	client := http.Client{
		Timeout: time.Second * timeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: time.Second * 30,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	return &Service{
		Client: client,
	}
}

//MakeRequest ...
func (s *Service) MakeRequest(ctx context.Context, payload RequestPayload) ([]byte, error) {
	req, err := http.NewRequest(payload.Method, payload.URL, bytes.NewReader(payload.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range payload.Headers {
		req.Header.Add(k, v)
	}

	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	//default success status
	if res.StatusCode > 210 {
		return body, fmt.Errorf("{status: %v, message: %v}", res.Status, string(body))
	}

	return body, nil
}

//PostHere ...
func (s *Service) PostHere(ctx context.Context, headers map[string]string, data string) {
	body := []byte(data)

	payload := RequestPayload{
		URL:     "https://posthere.io/f8c4-4160-b821",
		Method:  http.MethodPost,
		Headers: headers,
		Body:    body,
	}
	_, err := s.MakeRequest(ctx, payload)
	if err != nil {
		log.Println(err.Error())
	}
}
