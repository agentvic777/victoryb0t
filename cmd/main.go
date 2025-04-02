package main

import (
	"fmt"
	"log"
	"victorytg/config"
	"victorytg/victoryBot2"
)

const (
	StartCommand = "/start"
)

func main() {

	menu := victoryBot2.NewInlineKeyboardMarkup(
		victoryBot2.NewInlineKeyboardRow(
			victoryBot2.NewInlineKeyboardButtonData("🇺🇦Ukraine", "UA"),
		),
		victoryBot2.NewInlineKeyboardRow(
			victoryBot2.NewInlineKeyboardButtonData("🇦🇪United Arab Emirates", "AE"),
		),
		victoryBot2.NewInlineKeyboardRow(
			victoryBot2.NewInlineKeyboardButtonData("🇨🇳China", "CN"),
		),
		victoryBot2.NewInlineKeyboardRow(
			victoryBot2.NewInlineKeyboardButtonData("🇦🇩Moldova", "MD"),
		),
		victoryBot2.NewInlineKeyboardRow(
			victoryBot2.NewInlineKeyboardButtonData("🇦🇱Albania", "AL"),
		),
	)
	token, err := config.ExtractTelegramToken()
	if err != nil {
		log.Fatal(err)
	}
	currentConfigs := victoryBot2.Configs{
		Token: token,
	}
	for update := range currentConfigs.GetUpdates() {
		if update.Message != nil {
			if update.Message.MessageEntity != nil {
				log.Printf("Received update: %+v", update.Message.Text)
				switch update.Message.Text {
				case StartCommand:
					msg := victoryBot2.NewMessage(update.Message.Chat.ID, "Select country:")
					msg.ReplyMarkup = &menu
					currentConfigs.SendMessage(msg)
				}
			}
		}
		if update.CallbackQuery != nil {
			fmt.Println(update.CallbackQuery.Data)
		}
	}
}

//func Something() {
//	const (
//		StartCommand = "/start"
//	)
//	var (
//		update victoryBot.Update
//		token  = "7970273901:AAGdFpp-dubGGrYeFUUnoWt86MkVR9Tdm20"
//		offset int
//	)
//	updClient, err := victoryBot.InitClient(token, "/getUpdates")
//	sendClient, err := victoryBot.InitClient(token, "/sendMessage")
//	if err != nil {
//		return
//	}
//	updatesChannel, lastID, err := victoryBot.GetUpdatesChannel(updClient, offset)
//	if err != nil {
//		return
//	}
//	offset = lastID
//	for {
//		update = <-updatesChannel
//		if update.Message != nil && update.Message.IsCommand() {
//			switch update.Message.Text {
//			case StartCommand:
//				msg := victoryBot.NewMessage(update.Message.Chat.ID, "We're starting!")
//				//msg.ReplyMarkup = menu
//				_, err = sendClient.SendMessage(msg)
//				if err != nil {
//					return
//				}
//			default:
//				msg := victoryBot.NewMessage(update.Message.Chat.ID, "Other")
//				_, err = victoryBot.SendMessage(sendClient, msg)
//				if err != nil {
//					return
//				}
//			}
//		}
//	}
//}
