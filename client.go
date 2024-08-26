package cf_forbidden

import (
	"crypto/tls"
	"net/http"
)

func New(browser Browser, debug bool) *http.Client {
	tlsConfig := &tls.Config{
		CipherSuites:       cipherSuites[browser],
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS13,
		InsecureSkipVerify: debug,
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
}
