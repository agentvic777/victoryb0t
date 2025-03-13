package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	url        string
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) CreateURL(method string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
	}
	pathSegments := []string{"bot" + c.token, method}
	u.Path = path.Join(pathSegments...)
	return u.String()
}

type Response struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"` // Raw JSON to be unmarshaled into specific types
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

func (c *Client) Call(method string, params interface{}) (*Response, error) {
	var body io.Reader
	c.url = c.CreateURL(method)

	if params != nil {
		jsonParams, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonParams)
	} else {
		body = nil
	}
	req, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apiResp Response

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp, nil
}
