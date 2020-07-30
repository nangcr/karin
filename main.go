package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/catsworld/qq-bot-api"
	"github.com/dimiro1/banner"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	banner.Init(os.Stdout, true, true, bytes.NewBufferString(`  _  __                 _         
 | |/ /                (_)        
 | ' /    __ _   _ __   _   _ __  
 |  <    / _`+"`"+` | | '__| | | | '_ \ 
 | . \  | (_| | | |    | | | | | |
 |_|\_\  \__,_| |_|    |_| |_| |_|

`))

	if DEBUG {
		log.Println("The program is running in debug mode")
	}

	db := redis.NewClient(&redis.Options{
		Addr:     DBADDR,
		Password: DBPASSWORD,
		DB:       DBDB,
	})
	_, err := db.Ping(ctx).Result()
	checkError(err)

	log.Println("Database connected")

	botAPI, err := qqbotapi.NewBotAPI(CQTOKEN, CQAPI, CQSECRET)
	checkError(err)
	botAPI.Debug = DEBUG

	log.Println("Bot API connected")

	bot, err := NewBot(botAPI, db, ALLOWGROUP)
	checkError(err)

	log.Println("Controller initialized.")

	bot.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
