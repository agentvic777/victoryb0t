package victoryBot2

import (
	"net/url"
)

func CreateURL(baseURL string, params ...string) (string, error) {
	result, err := url.JoinPath(baseURL, params...)
	if err != nil {
		return "", err
	}
	return result, err
}
