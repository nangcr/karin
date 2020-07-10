package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
)

func (bot *Bot) handleRepeat(msg *qqbotapi.Message) {
	_, message := msg.Command()
	err := bot.SendMessages(msg.Chat.ID, msg.Chat.Type, strings.Join(message, ""))
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
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://service-kjcbcnmw-1254119946.gz.apigw.tencentcs.com/line", strings.NewReader("{\"history\":0}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", " https://kengxxiao.github.io/Kyouka/")
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	r := &clanApiLine{}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Panic(err)
	}

	message := cqcode.NewMessage()
	err = message.Append(&cqcode.Text{Text: "查询结果"})
	if err != nil {
		log.Panic(err)
	}

	for _, v := range r.Data {
		str := fmt.Sprintf("第%d名 %s 分数 %d", v.Rank, v.ClanName, v.Damage)
		err = message.Append(&cqcode.Text{Text: "\n" + str})
		if err != nil {
			log.Panic(err)
		}

		if v.Rank == 10000 {
			break
		}
	}

	err = bot.SendMessages(msg.Chat.ID, msg.Chat.Type, message)
	if err != nil {
		log.Panic(err)
	}
}
