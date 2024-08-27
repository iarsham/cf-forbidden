package cf_forbidden

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"net/http"
)

type M map[string]string

type Client struct {
	tls     cycletls.CycleTLS
	headers M
	ja3     string
	agent   string
}

func New() (*Client, error) {
	client := &Client{
		tls: cycletls.Init(),
	}
	return client, nil
}

func (c Client) Do(url string, opt cycletls.Options, method string) (cycletls.Response, error) {
	for k, v := range c.headers {
		opt.Headers[k] = v
	}
	if opt.UserAgent == "" {
		opt.UserAgent = c.agent
	}
	if opt.Ja3 == "" {
		opt.Ja3 = c.ja3
	}
	return c.tls.Do(url, opt, method)
}

func (c Client) Get(url string, headers M) (cycletls.Response, error) {
	opt := cycletls.Options{
		UserAgent: c.agent,
		Headers:   headers,
		Ja3:       c.ja3,
	}
	return c.Do(url, opt, http.MethodGet)
}

func (c Client) Post(url string, headers M, body string) (cycletls.Response, error) {
	options := cycletls.Options{
		Ja3:       c.ja3,
		Body:      body,
		UserAgent: c.agent,
		Headers:   headers,
	}
	return c.Do(url, options, http.MethodPost)
}
