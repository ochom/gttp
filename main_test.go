package gohttp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	s := NewHTTPService(30)
	res, err := s.Get(context.Background(), "https://google.com", nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
}

func TestService_Post(t *testing.T) {
	s := NewHTTPService(30)
	url := "https://posthere.io/f8c4-4160-b821"
	res, err := s.Post(context.Background(), url, map[string]string{"Authorization": "Basic hello world"}, []byte(`{"hello":"world"}`))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
}
