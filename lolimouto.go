package main

import (
	_ "embed"
	"log"
	"lolimouto/openweathermap"
	"math/rand"
	"net/http"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var gb, gbErr = tb.NewBot(tb.Settings{
	Token:  apiKey,
	Poller: &tb.LongPoller{Timeout: 1 * time.Second},
})

var httpClient = http.Client{
	Timeout: time.Second * 10,
}

var weatherClient = openweathermap.NewClient(owmApiKey)

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
