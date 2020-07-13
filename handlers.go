package main

import (
	"fmt"
	"log"
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

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
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
	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
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

	err = message.Append(&cqcode.Text{Text: "Pong!"})
	if err != nil {
		log.Panic(err)
	}

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleTimelineSave(msg *qqbotapi.Message) {
	tag := "timeline"
	cmd, _ := msg.Command()
	text := strings.TrimLeft(msg.Text, cmd)
	text = strings.TrimSpace(text)
	key := strings.Split(text, "\r\n")[0]
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

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleTimelineSearch(msg *qqbotapi.Message) {
	tag := "timeline"
	_, args := msg.Command()
	var key, result string
	var err error

	if args == nil {
		var data []string
		data, err = bot.searchData(tag)
		sort.Strings(data)
		result = strings.Join(data, "\n")
		result = strings.Replace(result, tag+":", "", -1)
	} else {
		key = args[0]
		result, err = bot.readData(tag, key)
		if err != nil {
			log.Panic(err)
		}
	}

	message := cqcode.NewMessage()
	err = message.Append(&cqcode.Text{Text: result})
	if err != nil {
		log.Panic(err)
	}

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleClanLine(msg *qqbotapi.Message) {
	line, updateTime, err := kyoka.GetLine()
	if err != nil {
		log.Panic(err)
	}

	message := cqcode.NewMessage()

	for _, v := range line {
		str := fmt.Sprintf("第%d名 %s 分数 %d", v.Rank, v.ClanName, v.Damage)
		err = message.Append(&cqcode.Text{Text: "\n" + str})
		if err != nil {
			log.Panic(err)
		}
	}

	err = message.Append(&cqcode.Text{Text: "\n更新时间 " + updateTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Panic(err)
	}

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleClanSearch(msg *qqbotapi.Message) {
	_, args := msg.Command()
	clans, updateTime, _, err := kyoka.GetByName(args[0], 0)
	if err != nil {
		log.Panic(err)
	}

	clans2, _, _, err := kyoka.GetByLeader(args[0], 0)
	if err != nil {
		log.Panic(err)
	}

	clans = append(clans, clans2...)

	message := cqcode.NewMessage()

	for _, v := range clans {
		str := fmt.Sprintf("第%d名 %s 会长 %s 分数 %d", v.Rank, v.ClanName, v.LeaderName, v.Damage)
		err = message.Append(&cqcode.Text{Text: "\n" + str})
		if err != nil {
			log.Panic(err)
		}
	}
	err = message.Append(&cqcode.Text{Text: "\n更新时间 " + updateTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Panic(err)
	}

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleRankSearch(msg *qqbotapi.Message) {
	_, args := msg.Command()
	rank, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panic(err)
	}

	clan, updateTime, err := kyoka.GetByRank(rank)
	if err != nil {
		log.Panic(err)
	}

	message := cqcode.NewMessage()

	str := fmt.Sprintf("第%d名 %s 会长 %s 分数 %d", clan.Rank, clan.ClanName, clan.LeaderName, clan.Damage)
	err = message.Append(&cqcode.Text{Text: "\n" + str})
	if err != nil {
		log.Panic(err)
	}

	err = message.Append(&cqcode.Text{Text: "\n更新时间 " + updateTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Panic(err)
	}

	err = bot.sendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}
