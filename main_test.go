package gttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	res, status, err := NewRequest("https://google.com", nil, nil).Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, nil, res)
}

func TestService_Post(t *testing.T) {
	url := "https://posthere.io/f8c4-4160-b821"
	headers := map[string]string{"Authorization": "Basic hello world"}
	body := []byte(`{"hello":"world"}`)
	res, status, err := NewRequest(url, headers, body).Post()
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, nil, res)
}

func TestWithTimeout(t *testing.T) {
	url := "https://posthere.io/f8c4-4160-b821"
	headers := map[string]string{"Authorization": "Basic hello world"}
	body := []byte(`{"hello":"world"}`)
	res, status, err := NewRequestWithTimeout(url, headers, body, 5).Post()
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, nil, res)
}

func TestNew(t *testing.T) {
	url := "https://posthere.io/f8c4-4160-b821"
	headers := map[string]string{"Authorization": "Basic hello world"}
	body := []byte(`{"hello":"world"}`)

	req := New()
	req.SetBody(body)
	req.SetHeaders(headers)
	req.SetURL(url)

	res, status, err := req.Post()
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, nil, res)
}
