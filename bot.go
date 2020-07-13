package main

import (
	"log"
	"strings"

	"github.com/catsworld/qq-bot-api"
	"github.com/go-redis/redis/v8"
	"github.com/nangcr/kyoka-tentacle"
)

var kyoka *kyokatentacle.API

func init() {
	var err error
	kyoka, err = kyokatentacle.NewAPI()
	if err != nil {
		checkError(err)
	}
}

type Bot struct {
	api        *qqbotapi.BotAPI
	db         *redis.Client
	allowGroup int64
	updates    <-chan qqbotapi.Update
}

func NewBot(api *qqbotapi.BotAPI, db *redis.Client, allowGroup int64) (bot *Bot, err error) {
	bot = &Bot{
		api:        api,
		db:         db,
		allowGroup: allowGroup,
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
	if msg != nil && (msg.Chat.IsPrivate() || (msg.Chat.IsGroup() && msg.Chat.ID == bot.allowGroup)) && strings.HasPrefix(msg.Text, "/") {
		log.Printf("[%s] %s", msg.From.String(), msg.Text)

		cmd, _ := msg.Command()
		if cmd == "/帮助" || cmd == "/help" {
			bot.handleHelp(msg)
		}
		if cmd == "/复读" {
			bot.handleRepeat(msg)
		}
		if cmd == "/ping" {
			bot.handlePing(msg)
		}
		if cmd == "/查线" {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Fatal: %+v\n", r)
					}
				}()
				bot.handleClanLine(msg)
			}()
		}
		if cmd == "/查公会" {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Fatal: %+v\n", r)
					}
				}()
				bot.handleClanSearch(msg)
			}()
		}
		if cmd == "/查排名" {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Fatal: %+v\n", r)
					}
				}()
				bot.handleRankSearch(msg)
			}()
		}
	}
}

func (bot *Bot) sendMessages(chatID int64, chatType string, message interface{}) (err error) {
	_, err = bot.api.SendMessage(chatID, chatType, message)

	return
}
