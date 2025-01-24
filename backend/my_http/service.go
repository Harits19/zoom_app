package my_http

import (
	"bytes"
	"fmt"
	"hello-world/access_token"
	"io"
	"net/http"
)

func StringToBuffer(value string) *bytes.Buffer {
	body := []byte(value)

	return bytes.NewBuffer(body)
}

func Request(method string, url string, body io.Reader, successStatusCode int) (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)

	if err != nil {

		return nil, err
	}

	access_token.SetAccessToken(request)

	response, err := client.Do(request)

	if err != nil {

		return nil, err
	}

	if response.StatusCode != successStatusCode {

		return nil, (fmt.Errorf("failed to create meeting with error code %d", response.StatusCode))
	}

	return response, nil
}
