package main

import (
	"strconv"
	"strings"

	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
)

func (bot *Bot) handleFuDu(msg *qqbotapi.Message) {
	_, message := msg.Command()
	err := bot.SendMessages(msg.Chat.ID, msg.Chat.Type, strings.Join(message, ""))
	checkError(err)
}

func (bot *Bot) handlePing(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	err := message.Append(&cqcode.At{QQ: strconv.FormatInt(msg.From.ID, 10)})
	checkError(err)

	err = message.Append(&cqcode.Text{Text: "Pong!"})
	checkError(err)

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
	checkError(err)
}
