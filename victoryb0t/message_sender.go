package victoryBot2

import "encoding/json"

type SendMessageRequest struct {
	ChatID      interface{}           `json:"chat_id"`
	Text        string                `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MsgResponse struct {
	OK          bool    `json:"ok"`
	Result      Message `json:"result,omitempty"`
	ErrorCode   int     `json:"error_code,omitempty"`
	Description string  `json:"description,omitempty"`
}

func NewMessage(chatID int64, text string) *SendMessageRequest {
	return &SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}
}

func (c *Configs) SendMessage(params *SendMessageRequest) (*Message, error) {
	var (
		serverResp MsgResponse
		url        string
		err        error
		byteData   []byte
	)

	url, err = CreateURL("https://api.telegram.org", "bot"+c.Token, "sendMessage")
	requester := APIRequester{
		HttpMethod:   "POST",
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

	return &serverResp.Result, nil
}
