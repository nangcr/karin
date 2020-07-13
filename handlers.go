package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
)

func (bot *Bot) handleHelp(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	err := message.Append(&cqcode.Text{Text: HELPSTRING})
	if err != nil {
		log.Panic(err)
	}

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *Bot) handleRepeat(msg *qqbotapi.Message) {
	message := cqcode.NewMessage()
	_, args := msg.Command()
	err := message.Append(&cqcode.Text{Text: strings.Join(args, "")})
	if err != nil {
		log.Panic(err)
	}
	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
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

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
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
	err = message.Append(&cqcode.Text{Text: "查询结果"})
	if err != nil {
		log.Panic(err)
	}

	for _, v := range line {
		str := fmt.Sprintf("第%d名 %s 分数 %d", v.Rank, v.ClanName, v.Damage)
		err = message.Append(&cqcode.Text{Text: "\n" + str})
		if err != nil {
			log.Panic(err)
		}

		if v.Rank == 10000 {
			err := message.Append(&cqcode.Text{Text: "\n更新时间 " + updateTime.Format("2006-01-02 15:04:05")})
			if err != nil {
				log.Panic(err)
			}
			break
		}
	}

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
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
	err = message.Append(&cqcode.Text{Text: "查询结果"})
	if err != nil {
		log.Panic(err)
	}

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

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
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
	err = message.Append(&cqcode.Text{Text: "查询结果"})
	if err != nil {
		log.Panic(err)
	}

	str := fmt.Sprintf("第%d名 %s 会长 %s 分数 %d", clan.Rank, clan.ClanName, clan.LeaderName, clan.Damage)
	err = message.Append(&cqcode.Text{Text: "\n" + str})
	if err != nil {
		log.Panic(err)
	}

	err = message.Append(&cqcode.Text{Text: "\n更新时间 " + updateTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Panic(err)
	}

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}
