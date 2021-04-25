package main

import (
	"log"

	"github.com/catsworld/qq-bot-api"
	"github.com/go-redis/redis/v8"
)

var skill map[string]int

func init() {
	skill = map[string]int{"闪耀": 300, "死炎法": 290, "煞星": 2000, "法令": 400, "苦难之心": 900, "崩石": 100000}
}

type Bot struct {
	api        *qqbotapi.BotAPI
	db         *redis.Client
	allowGroup []int64
	updates    <-chan qqbotapi.Update
}

func NewBot(api *qqbotapi.BotAPI, db *redis.Client, allowGroup []int64) (bot *Bot, err error) {
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
	if msg != nil && (msg.Chat.IsPrivate() || (msg.Chat.IsGroup() && inIntSlice(bot.allowGroup, msg.Chat.ID))) {
		log.Printf("[%s] %s", msg.From.String(), msg.Text)

		cmd, _ := msg.Command()
		if cmd == "帮助" || cmd == "help" {
			bot.handleHelp(msg)
		}
		if cmd == "复读" {
			bot.handleRepeat(msg)
		}
		if cmd == "ping" {
			bot.handlePing(msg)
		}
		if cmd == "存轴" {
			bot.handleTimelineSave(msg)
		}
		if cmd == "查轴" {
			bot.handleTimelineSearch(msg)
		}
		if cmd == "删轴" {
			bot.handleTimelineDelete(msg)
		}
		if cmd == "大西瓜" {
			bot.handleReplyString(msg, "http://www.wesane.com/game/654")
		}
		if cmd == "大百合" {
			bot.handleReplyString(msg, "http://192.144.170.228:12450/")
		}
		if cmd == "魔塔" {
			bot.handleReplyString(msg, "https://h5mota.com/")
		}
		if potency, ok := skill[cmd]; ok {
			bot.handleDamage(msg, potency)
		}
	}
}

func (bot *Bot) saveData(tag, key string, value interface{}) (err error) {
	err = bot.db.Set(ctx, tag+":"+key, value, 0).Err()
	return
}

func (bot *Bot) readData(tag, key string) (result string, err error) {
	result, err = bot.db.Get(ctx, tag+":"+key).Result()
	return
}

func (bot *Bot) searchData(tag string) (result []string, err error) {
	result, err = bot.db.Keys(ctx, tag+":*").Result()
	return
}

func (bot *Bot) deleteData(tag, key string) (result int64, err error) {
	result, err = bot.db.Del(ctx, tag+":"+key).Result()
	return
}

func inIntSlice(haystack []int64, needle int64) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}

	return false
}
