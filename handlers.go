package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"math/rand"
	"strconv"
	"strings"
)

func shoutHandler(m *tb.Message) {
	msg := strings.ToUpper(strings.Replace(m.Text, "/shout ", "", 1))

	if msg == "/SHOUT" {
		return
	}

	split := strings.Split(msg, "")
	res := strings.Join(split, " ")

	for i, s := range split[1:] {
		res = res + "\n" + s + fmt.Sprintf("%"+strconv.Itoa((i+1)*2)+"s", s)
	}

	_, _ = gb.Send(m.Chat, fmt.Sprintf("`%s`", res), &tb.SendOptions{
		ParseMode: tb.ModeMarkdown,
	})
}

func textHandler(m *tb.Message) {
	if strings.HasPrefix(m.Text, "s/") && m.ReplyTo != nil {
		split := strings.Split(m.Text, "/")

		if len(split) < 3 {
			return
		}
		replyTo := m.ReplyTo
		replyText := "`Did you mean:` \n" + strings.Replace(replyTo.Text, split[1], split[2], -1)

		_, _ = gb.Reply(replyTo, replyText, &tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		})

		return
	}

	textLower := strings.ToLower(m.Text)
	if strings.HasSuffix(textLower, " y/n") || strings.HasPrefix(textLower, "y/n") {
		responses := []string{"Yes", "No"}
		i := rand.Intn(len(responses))

		_, _ = gb.Send(m.Chat, responses[i])
	}
}
