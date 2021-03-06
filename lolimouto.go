package main

import (
	_ "embed"
	"fmt"
	"log"
	"lolimouto/openweathermap"
	"math/rand"
	"net/http"
	"strings"
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
	gb.Handle("/hot", hotHandler)
	gb.Handle("/loli", loliHandler)
	gb.Handle("/weather", weatherHandler)
	gb.Handle("/eval", evalHandler)
	gb.Handle("/roll", rollHandler)
	gb.Handle(tb.OnText, textHandler)
}

func somethingWentWrong(m *tb.Message) {
	_, _ = gb.Send(m.Sender, "Something went wrong. Please try again later.")
}

func implodeIntSlice(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
