package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func Bot(w http.ResponseWriter, r *http.Request) {
	//token := os.Getenv("BOT_TOKEN")
	//bot, error := tgbotapi.NewBotAPI(token)

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	log.Printf("[%s] %s R29", update.Message.From.UserName, update.Message.Text)

	text := "hej hej"
	data := Response{Msg: text,
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID}

	msg, _ := json.Marshal(data)
	log.Printf("Response %s", string(msg))
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(msg))

	//if update.Message.IsCommand() {
	//	text := "hej hej"
	//	switch update.Message.Command() {
	//	case "hola":
	//		text = "someeeeeee"
	//	case "hej":
	//		text = fmt.Sprintf("→ something")
	//	default:
	//		text = "fallback greeting"
	//	}
	//	data := Response{Msg: text,
	//		Method: "sendMessage",
	//		ChatID: update.Message.Chat.ID}
	//
	//	msg, _ := json.Marshal(data)
	//	log.Printf("Response %s", string(msg))
	//	w.Header().Add("Content-Type", "application/json")
	//	fmt.Fprintf(w, string(msg))
	//}
}
