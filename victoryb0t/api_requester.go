package victoryBot2

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type APIRequester struct {
	HttpMethod   string
	Url          string
	RequestParam interface{}
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func (r *APIRequester) NewHTTPRequest() ([]byte, error) {
	var requestBodyReader io.Reader
	if r.RequestParam != nil {
		jsonData, err := json.Marshal(r.RequestParam)
		if err != nil {
			return nil, err
		}
		requestBodyReader = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest(r.HttpMethod, r.Url, requestBodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := NewHTTPClient()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
