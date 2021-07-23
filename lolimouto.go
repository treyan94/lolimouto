package main

import (
	_ "embed"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var apiKey = func() (key string) {
	key = os.Getenv("LOLIMOUTO_BOT_KEY")

	if args := os.Args[1:]; len(args) != 0 {
		key = args[0]
	}

	if key == "" {
		log.Fatal("provide telegram bot api key as first argument")
	}

	return key
}()

var gb, gbErr = tb.NewBot(tb.Settings{
	Token:  apiKey,
	Poller: &tb.LongPoller{Timeout: 1 * time.Second},
})

var httpClient = http.Client{
	Timeout: time.Second * 10,
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if gbErr != nil {
		log.Fatal(gbErr)
	}

	registerHandlers()

	gb.Start()
}

func registerHandlers() {
	gb.Handle("/shout", shoutHandler)
	gb.Handle("/pathetic", patheticHandler)
	gb.Handle("/blasphemy", blasphemyHandler)
	gb.Handle("/loli", loliHandler)
	gb.Handle(tb.OnText, textHandler)
}
