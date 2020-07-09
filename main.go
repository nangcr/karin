package main

import (
	"flag"
	"log"

	"github.com/catsworld/qq-bot-api"
)

func main() {
	token := flag.String("token", TOKEN, "CoolQ token")
	api := flag.String("api", API, "CoolQ address and port")
	secret := flag.String("secret", SECRET, "CoolQ http secret")
	flag.Parse()

	botapi, err := qqbotapi.NewBotAPI(*token, *api, *secret)
	checkError(err)
	botapi.Debug = DEBUG

	log.Println("Bot API connected")

	bot, err := NewBot(botapi)
	checkError(err)

	log.Println("Controller initialized.")

	bot.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
