package victoryBot2

import "log"

type CommandHandler func(update Update) (*SendMessageRequest, error)
type CallbackHandler func(update Update) (*SendMessageRequest, error)

type Bot struct {
	Config          *Configs
	Commands        map[string]CommandHandler
	CallbackHandler CallbackHandler
}

func InitBot(configs *Configs) *Bot {
	return &Bot{
		Config:   configs,
		Commands: make(map[string]CommandHandler),
	}
}

func (bot *Bot) RegisterCommand(command string, handler CommandHandler) {
	bot.Commands[command] = handler
}

func (bot *Bot) RegisterCallback(callback CallbackHandler) {
	bot.CallbackHandler = callback
}

func (bot *Bot) Run() {
	updates := bot.Config.GetUpdates()
	for update := range updates {
		if update.Message != nil {
			if update.Message.MessageEntity != nil {
				command := update.Message.Text
				if handler, ok := bot.Commands[command]; ok {
					response, err := handler(update)
					if err != nil {
						log.Printf("Error while executing command %s: %s", command, err)
						continue
					}
					if response != nil {
						_, err := bot.Config.SendMessage(response)
						if err != nil {
							log.Printf("Error sending message %s: %s", response, err)
							continue
						}
					}
				}
			}
		}
		if update.CallbackQuery != nil && bot.CallbackHandler != nil {
			response, err := bot.CallbackHandler(update)
			if err != nil {
				log.Printf("Error while executing callback: %s", err)
				continue
			}
			if response != nil {
				_, err := bot.Config.SendMessage(response)
				if err != nil {
					log.Printf("Error sending message %s: %s", response, err)
					continue
				}
			}
		}
	}
}
