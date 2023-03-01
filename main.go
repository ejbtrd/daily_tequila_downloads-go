package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("Download stats as of %s in last 24 hours:", time)

	const CHAT_ID int64 = -1001791062372

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	msg := tgbotapi.NewMessage(CHAT_ID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
