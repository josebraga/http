package httptls

import (
	"crypto/tls"
	"net/http"
)

func NewTLSClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func NewTLSServer() *http.Server {
	return &http.Server{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}
