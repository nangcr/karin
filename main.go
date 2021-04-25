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
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Database connected")

	botAPI, err := qqbotapi.NewBotAPI(CQTOKEN, CQAPI, CQSECRET)
	if err != nil {
		log.Fatalln(err)
	}
	botAPI.Debug = DEBUG

	log.Println("Bot API connected")

	bot, err := NewBot(botAPI, db, []int64{149380170,809653356})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Controller initialized.")

	bot.Run()
}