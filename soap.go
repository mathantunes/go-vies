package vies

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type soap struct {
	httpClient *http.Client
}

func newSoap() *soap {
	return &soap{
		httpClient: &http.Client{},
	}
}

type soapRequester interface {
	MakeRequest(endpoint, action string, payload []byte) (io.ReadCloser, error)
}

func (s *soap) MakeRequest(endpoint, action string, payload []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodPost, VIESEndpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", fmt.Sprintf("urn:%s", action))
	client := s.httpClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
