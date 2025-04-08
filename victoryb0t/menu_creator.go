package victoryBot2

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton

	keyboard = append(keyboard, rows...)

	return InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton

	row = append(row, buttons...)

	return row
}

func NewInlineKeyboardButtonData(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}
