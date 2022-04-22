package api

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/charlie-haley/vyconfigure/pkg/options"
)

type Client struct {
	Options    *options.Options
	httpClient *http.Client
}

func CreateClient(o *options.Options) (*Client, error) {
	t := http.DefaultTransport.(*http.Transport).Clone()
	if o.Insecure {
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	httpClient := &http.Client{Transport: t, Timeout: time.Duration(10) * time.Second}

	client := &Client{
		Options:    o,
		httpClient: httpClient,
	}

	return client, nil
}
