package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	AddNewWord = "AddNewWord"
	GetTest    = "GetTest"
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AddNewWord),
		tgbotapi.NewKeyboardButton(GetTest),
	),
)

func main() {

	f, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TokenTg"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var messageText string = update.Message.Text
		switch update.Message.Text {
		case AddNewWord:
			messageText = "Enter the word!"
		case GetTest:
			messageText = "Test starting!"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		msg.ReplyMarkup = mainKeyboard
		bot.Send(msg)
	}
}
