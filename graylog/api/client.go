package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiBaseURL string
	token      string
	username   string
	verify     bool
	httpClient *http.Client
}

// NewClient returns a new client that can access the Graylog API
func NewClient(url string, token string, username string, verify bool) *Client {
	var tr *http.Transport

	if verify {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}

	return &Client{
		apiBaseURL: fmt.Sprintf("%s/api", url),
		token:      token,
		username:   username,
		verify:     verify,
		httpClient: c,
	}
}
