package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
)

func (bot *Bot) handleHelp(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	err := message.Append(&cqcode.Text{Text: HELPSTRING})
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleRepeat(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	cmd, _ := msg.Command()
	text := strings.TrimLeft(msg.Text, cmd)
	text = strings.TrimSpace(text)
	err := message.Append(&cqcode.Text{Text: text})
	if err != nil {
		log.Panic(err)
	}
	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handlePing(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	err := message.Append(&cqcode.At{QQ: strconv.FormatInt(msg.From.ID, 10)})
	if err != nil {
		log.Panic(err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(10) == 0 {
		err = message.Append(&cqcode.Text{Text: "pia!"})
	} else {
		err = message.Append(&cqcode.Text{Text: "pong!"})
	}
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleTimelineSave(msg *qqbotapi.Message) {
	tag := "timeline"
	cmd, _ := msg.Command()
	text := strings.TrimLeft(msg.Text, cmd)
	text = strings.TrimSpace(text)
	key := strings.Split(text, "\n")[0]
	key = strings.TrimSpace(key)
	suffix := fmt.Sprintf("\n由%s上传，上传时间 %s", msg.From.Name(), time.Now().Format("2006-01-02 15:04:05"))

	err := bot.saveData(tag, key, text+suffix)
	if err != nil {
		log.Panic(err)
	}

	message := cqcode.NewMessage()
	err = message.Append(&cqcode.Text{Text: key + "已保存"})
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleTimelineSearch(msg *qqbotapi.Message) {
	tag := "timeline"
	cmd, _ := msg.Command()
	text := strings.TrimLeft(msg.Text, cmd)
	text = strings.TrimSpace(text)
	key := strings.Split(text, "\n")[0]
	key = strings.TrimSpace(key)
	var result string
	var err error

	if key == "" {
		var data []string
		data, err := bot.searchData(tag)
		if err != nil {
			log.Panic(err)
		}

		sort.Strings(data)
		result = strings.Join(data, "\n")
		result = strings.Replace(result, tag+":", "", -1)
		result = "为您找到以下轴\n" + result
	} else {
		result, err = bot.readData(tag, key)
		if err != nil {
			result = "未找到 " + key
		}
	}

	message := cqcode.NewMessage()
	err = message.Append(&cqcode.Text{Text: result})
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleTimelineDelete(msg *qqbotapi.Message) {
	tag := "timeline"
	cmd, _ := msg.Command()
	text := strings.TrimLeft(msg.Text, cmd)
	text = strings.TrimSpace(text)
	key := strings.Split(text, "\n")[0]
	key = strings.TrimSpace(key)
	var result string

	if key == "" {
		data, err := bot.searchData(tag)
		if err != nil {
			log.Panic(err)
		}

		sort.Strings(data)
		result = strings.Join(data, "\n")
		result = strings.Replace(result, tag+":", "", -1)
	} else {
		ok, err := bot.deleteData(tag, key)
		if err != nil {
			log.Panic(err)
		}

		if ok == 1 {
			result = key + " 删除成功"
		} else {
			result = "未找到 " + key
		}
	}

	message := cqcode.NewMessage()
	err := message.Append(&cqcode.Text{Text: result})
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleReplyString(msg *qqbotapi.Message, reply string) {
	message := cqcode.NewMessage()
	err := message.Append(&cqcode.Text{Text: reply})
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleDamage(msg *qqbotapi.Message, potency int) {
	message := cqcode.NewMessage()
	crit := false
	direct := false

	if potency>2000 && msg.From.ID != 2787019693 && msg.From.ID != 1658873149{
		message.Append(&cqcode.Text{Text: "不许崩石"})
		_, err := bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
		if err != nil {
			log.Panic(err)
		}
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	damage := float64(potency) * 83.3
	if r.Intn(100) < 34 {
		crit = true
		damage *= 1.63
	}

	if r.Intn(100) < 13 {
		direct = true
		damage *= 1.25
	}
	damage += damage * float64(r.Intn(50)-24) / 1000

	if crit && direct {
		message.Append(&cqcode.Text{Text: fmt.Sprintf("直击加暴击！木人受到了%d点伤害", int(damage))})
	} else if crit {
		message.Append(&cqcode.Text{Text: fmt.Sprintf("暴击！木人受到了%d点伤害", int(damage))})
	} else if direct {
		message.Append(&cqcode.Text{Text: fmt.Sprintf("直击！木人受到了%d点伤害", int(damage))})
	} else {
		message.Append(&cqcode.Text{Text: fmt.Sprintf("木人受到了%d点伤害", int(damage))})
	}

	_, err := bot.api.SendMessage(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}
