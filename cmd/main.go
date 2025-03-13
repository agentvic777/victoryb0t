package main

import (
	"fmt"
	"log"
	"victorytg/config"
)

func main() {
	telegramToken, err := config.ExtractTelegramToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Telegram Token: %s\n", telegramToken)

}
