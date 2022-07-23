package httpClient

import (
	"fmt"
	"net/http"
	"time"
)

const (
	httpTimeout = 15 * time.Second
)

func MakeHttpRequest(method string, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	client := http.Client{
		Timeout: httpTimeout,
	}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return response, fmt.Errorf("%s request to %s failed: %d", method, url, response.StatusCode)
	}
	return response, nil
}
