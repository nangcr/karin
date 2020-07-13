package main

import (
	"context"
	"log"

	"github.com/catsworld/qq-bot-api"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
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
