package cf_forbidden

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"net/http"
)

type M map[string]string

type Client struct {
	tls       cycletls.CycleTLS
	Headers   M
	Ja3       string
	UserAgent string
}

func New() (*Client, error) {
	client := &Client{
		tls: cycletls.Init(),
	}
	return client, nil
}

func (c Client) Do(url string, opt cycletls.Options, method string) (cycletls.Response, error) {
	for k, v := range c.Headers {
		opt.Headers[k] = v
	}
	if opt.Ja3 == "" {
		opt.Ja3 = c.Ja3
	}
	if opt.UserAgent == "" {
		opt.UserAgent = c.UserAgent
	}
	return c.tls.Do(url, opt, method)
}

func (c Client) Get(url string, headers M) (cycletls.Response, error) {
	opt := cycletls.Options{
		Headers:   headers,
		Ja3:       c.Ja3,
		UserAgent: c.UserAgent,
	}
	return c.Do(url, opt, http.MethodGet)
}

func (c Client) Post(url string, headers M, body string) (cycletls.Response, error) {
	options := cycletls.Options{
		Headers:   headers,
		Ja3:       c.Ja3,
		UserAgent: c.UserAgent,
		Body:      body,
	}
	return c.Do(url, options, http.MethodPost)
}
