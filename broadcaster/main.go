package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

type todo struct {
	ID			int			`json:"id"`
	Name		string		`json:"name" pg:",notnull"`
	Done		bool		`json:"done"`
}

var chatID int64 = 846849544

func main() {
	// Telegram
	godotenv.Load()
	botToken := os.Getenv("TG_BOT_TOKEN")
	if botToken == "" {
		log.Println("No bot token provided")
		os.Exit(1);
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Println("Authorized on account", bot.Self.UserName)

	// Nats
	natsAddr := os.Getenv("NATS")
	if natsAddr == "" {
		log.Println("No NATS address provided")
		os.Exit(1);
	}
	conn, err := nats.Connect(natsAddr)
	if err != nil {
		panic(err)
	}
	
	inputChan := make(chan *nats.Msg, 64)
	todoSub, err := conn.ChanQueueSubscribe("todo", "todo.broadcasters", inputChan)
	if err != nil {
		panic(err)
	}
	defer func() {
		todoSub.Unsubscribe()
		close(inputChan)
	}()

	func() {
		for todoMsg := range inputChan {
			msgText := string(todoMsg.Data[:])
			msg := tgbotapi.NewMessage(chatID, msgText)
			
			sentMsg, err := bot.Send(msg)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println("Sent message:", sentMsg.Text)
		}
	}()
}