package victoryBot2

type Configs struct {
	Token string
}

type Update struct {
	Offset        int            `json:"offset,omitempty"`
	UpdateID      int            `json:"update_id,omitempty"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type CallbackQuery struct {
	ID      string   `json:"id"`
	From    User     `json:"from"`
	Data    string   `json:"data,omitempty"`
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID     int                   `json:"message_id,omitempty"`
	Date          int                   `json:"date,omitempty"`
	From          User                  `json:"from,omitempty"`
	Text          string                `json:"text"`
	MessageEntity *[]MessageEntity      `json:"entities,omitempty"`
	Chat          Chat                  `json:"chat"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type User struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type MessageEntity struct {
	Type string `json:"type"`
}

type Chat struct {
	ID int64 `json:"id"`
}
