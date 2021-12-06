package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//RequestPayload is the format required to make request
type RequestPayload struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte

	//ResponseCode if not provided, status 200 is used
	ResponseCode *int
}

//HTTPService blueprint to the available functions
type HTTPService interface {
	MakeRequest(ctx context.Context, payload RequestPayload) ([]byte, error)
}

//Service the controller struct for external access
type Service struct {
	Client http.Client
}

//NewHTTPService create a new instance of HTTPService
func NewHTTPService(timeout time.Duration) HTTPService {
	client := http.Client{
		Timeout: time.Second * timeout,
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

	//return  based on user expected status code
	if payload.ResponseCode != nil {
		if res.StatusCode == *payload.ResponseCode {
			return body, nil
		}
		return body, fmt.Errorf("{status: %v, message: %v}", res.Status, string(body))
	}

	//default success status is 200 Ok
	if res.StatusCode != http.StatusOK {
		return body, fmt.Errorf("{status: %v, message: %v}", res.Status, string(body))
	}

	return body, nil
}
