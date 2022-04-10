package gohttp

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	s := New(time.Second * 30)
	status, res, err := s.Get(context.Background(), "https://google.com", nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
	assert.Equal(t, 200, status, "response status should be 200")
}

func TestService_Post(t *testing.T) {
	s := New(time.Second * 30)
	url := "https://posthere.io/f8c4-4160-b821"
	status, res, err := s.Post(context.Background(), url, map[string]string{"Authorization": "Basic hello world"}, []byte(`{"hello":"world"}`))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
	assert.Equal(t, 200, status, "response status should be 200")
}
