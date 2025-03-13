package telegram

import (
	"encoding/json"
	"log"
)

func (c *Client) GetUpdates() []Update {
	var params interface{}
	res, err := c.Call("/getUpdates", params)
	if err != nil {
		log.Fatal(err)
	}
	var updates []Update
	if err := json.Unmarshal(res.Result, &updates); err != nil {
		log.Fatalf("Failed to parse result: %v", err)
	}
	return updates
}
