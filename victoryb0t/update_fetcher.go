package victoryBot2

import (
	"encoding/json"
	"fmt"
)

type UpdResponse struct {
	OK          bool     `json:"ok"`
	Result      []Update `json:"result,omitempty"`
	ErrorCode   int      `json:"error_code,omitempty"`
	Description string   `json:"description,omitempty"`
}

type UpdatesChannel struct {
	LastUpdateID int
	Updates      chan Update
}

func (c *Configs) GetUpdates() chan Update {
	lastUpdateID := 0 // Start from the beginning or load from persistent storage
	updatesChan := InitUpdatesChannel(lastUpdateID)
	processUpdates(updatesChan, "POST", c.Token, "getUpdates")
	return updatesChan.Updates
}

func InitUpdatesChannel(lastUpdateID int) *UpdatesChannel {
	ch := make(chan Update, 100)
	return &UpdatesChannel{
		LastUpdateID: lastUpdateID,
		Updates:      ch,
	}
}

func processUpdates(ch *UpdatesChannel, httpMethod, token, method string) UpdatesChannel {
	go func() {
		defer close(ch.Updates)
		for {
			updates, err := fetchUpdates(ch.LastUpdateID+1, httpMethod, token, method)
			if err != nil {
				fmt.Printf("Error fetching updates: %v\n", err)
				continue
			}
			for _, update := range updates {
				if update.UpdateID > ch.LastUpdateID {
					ch.LastUpdateID = update.UpdateID
					ch.Updates <- update

				}
			}
		}
	}()
	return *ch
}

func fetchUpdates(offset int, httpMethod, token, method string) ([]Update, error) {
	var (
		serverResp UpdResponse
		params     = map[string]interface{}{
			"offset": offset,
		}
		url      string
		err      error
		byteData []byte
	)

	url, err = CreateURL("https://api.telegram.org", "bot"+token, method)
	requester := APIRequester{
		HttpMethod:   httpMethod,
		Url:          url,
		RequestParam: params,
	}
	if err != nil {
		return nil, err
	}

	byteData, err = requester.NewHTTPRequest()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteData, &serverResp)
	if err != nil {
		return nil, err
	}

	return serverResp.Result, err
}
