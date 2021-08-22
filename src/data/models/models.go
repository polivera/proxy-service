package models

import (
	"encoding/json"
	"io"
	"net/http"
)

// ProxyRequest Save all request that pass through the proxy
type ProxyRequest struct {
	ID              string
	Method          string
	Host            string
	URI             string
	Headers         string
	Body            string
	ResponseHeaders string
	ResponseBody    string
	ResponseStatus  uint
	Status          uint8
}

// RequestConfig Save where to proxy the incoming requests
type RequestConfig struct {
	ID            uint
	Source        string
	Target        string
	OriginHeader  string
	RefererHeader string
}

func (pr *ProxyRequest) LoadRequest(r *http.Request) error {
	var (
		headers, body []byte
		err           error
	)

	if headers, err = json.Marshal(r.Header); err != nil {
		return err
	}
	if body, err = io.ReadAll(r.Body); err != nil {
		return err
	}

	pr.Method = r.Method
	pr.Host = r.URL.Host
	pr.URI = r.URL.RawQuery
	pr.Headers = string(headers)
	pr.Body = string(body)
	return nil
}
