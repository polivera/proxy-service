package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

//todo join this methods

func ReadAndResetRequestBody(rq *http.Request) (string, error) {
	content, errRead := io.ReadAll(rq.Body)
	_ = rq.Body.Close()
	if errRead != nil {
		return "", errRead
	}
	rq.Body = io.NopCloser(bytes.NewBuffer(content))
	return string(content), nil
}

func ReadAndResetResponseBody(rs *http.Response) (string, error) {
	content, errRead := io.ReadAll(rs.Body)
	_ = rs.Body.Close()
	if errRead != nil {
		return "", errRead
	}
	rs.Body = io.NopCloser(bytes.NewBuffer(content))
	return string(content), nil
}

func GetHeadersAsJson(headers http.Header) (string, error) {
	jsonHeaders, err := json.Marshal(headers)
	return string(jsonHeaders), err
}
