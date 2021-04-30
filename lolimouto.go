package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  apiKey,
		Poller: &tb.LongPoller{Timeout: 1 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/shout", func(m *tb.Message) {
		msg := strings.ToUpper(strings.Replace(m.Text, "/shout ", "", 1))

		split := strings.Split(msg, "")
		res := ""

		for i, s := range split {
			if i == 0 {
				res = strings.Join(split, " ")
				continue
			}

			res = res + "\n" + s + fmt.Sprintf("%"+strconv.Itoa(i*2)+"s", s)
		}

		_, _ = b.Send(m.Chat, fmt.Sprintf("`%s`", res), &tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		})
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if strings.HasPrefix(m.Text, "s/") && m.ReplyTo != nil {
			split := strings.Split(m.Text, "/")

			if len(split) < 3 {
				return
			}
			replyTo := m.ReplyTo
			replyText := strings.Replace(replyTo.Text, split[1], split[2], -1)
			replyText = "Did you mean: \n" + replyText

			_, _ = b.Reply(replyTo, replyText)

			return
		}

		textLower := strings.ToLower(m.Text)
		if strings.HasSuffix(textLower, " y/n") || strings.HasPrefix(textLower, "y/n") {
			responses := []string{"Yes", "No"}
			i := rand.Intn(len(responses))

			_, _ = b.Send(m.Chat, responses[i])
		}
	})

	b.Start()
}
