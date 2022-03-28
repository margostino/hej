package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Job(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	botToken := os.Getenv("BOT_TOKEN")
	jobSecret := os.Getenv("HEJ_JOB_SECRET")
	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	log.Printf("logging start")
	fmt.Println("printing start")

	if r.Method == "POST" {
		requestSecret := r.Header.Get("Authorization")

		if requestSecret == fmt.Sprintf("Bearer %s", jobSecret) {
			bot, error := tgbotapi.NewBotAPI(botToken)

			if error != nil {
				panic(error)
			}

			message := tgbotapi.NewMessage(chatId, "hej hej JOB!")
			bot.Send(message)

			msg := tgbotapi.NewMessage(chatId, "hej")
			msg.ReplyMarkup = nil
			//msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

		}
	}

	log.Printf("logging end")
	fmt.Println("printing end")

}
