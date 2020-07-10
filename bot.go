package main

import (
	"log"
	"strings"

	"github.com/catsworld/qq-bot-api"
)

type Bot struct {
	api     *qqbotapi.BotAPI
	updates <-chan qqbotapi.Update
}

func NewBot(api *qqbotapi.BotAPI) (bot *Bot, err error) {
	bot = &Bot{
		api: api,
	}

	u := qqbotapi.NewUpdate(0)
	u.PreloadUserInfo = true
	u.Timeout = 60
	bot.updates, err = api.GetUpdatesChan(u)

	return
}

func (bot *Bot) Run() {
	for update := range bot.updates {
		bot.processUpdate(&update)
	}
}

func (bot *Bot) processUpdate(update *qqbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Fatal: %+v\n", r)
		}
	}()

	msg := update.Message
	if msg != nil && strings.HasPrefix(msg.Text, "/") {
		log.Printf("[%s] %s", msg.From.String(), msg.Text)

		cmd, _ := msg.Command()
		if cmd == "/复读" {
			bot.handleRepeat(msg)
		}
		if cmd == "/ping" {
			bot.handlePing(msg)
		}
		if cmd == "/查线" {
			bot.handleClanLine(msg)
		}
	}
}

func (bot *Bot) SendMessages(chatId int64, chatType string, message interface{}) (err error) {
	_, err = bot.api.SendMessage(chatId, chatType, message)

	return
}
