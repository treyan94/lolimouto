package main

import (
	"bytes"
	_ "embed"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"math/rand"
	"regexp"
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
	if (strings.HasPrefix(m.Text, "s/") || strings.HasPrefix(m.Text, "/s/")) && m.ReplyTo != nil {
		replaceText(m)

		return
	}

	textLower := strings.ToLower(m.Text)
	if strings.HasSuffix(textLower, " y/n") ||
		strings.HasPrefix(textLower, "y/n") ||
		strings.HasSuffix(textLower, " y/n?") {
		responses := []string{"Yes", "No"}
		i := rand.Intn(len(responses))

		_, _ = gb.Send(m.Chat, responses[i])
	}
}

func replaceText(m *tb.Message) {
	split := strings.Split(m.Text, "/")

	if len(split) < 3 {
		return
	}

	if split[0] == "" {
		split = split[1:]
	}

	replyTo := m.ReplyTo
	replyText := replyTo.Text

	if replyText == "" {
		replyText = replyTo.Caption
	}

	if replyText == "" {
		return
	}

	mustCompile := regexp.MustCompile(split[1])
	replyMessage := "`Did you mean:` \n" + mustCompile.ReplaceAllString(replyText, strings.Join(split[2:], "/"))

	_, _ = gb.Reply(replyTo, replyMessage, &tb.SendOptions{ParseMode: tb.ModeMarkdown})
}

//go:embed pathetic.png
var patheticImg []byte

func patheticHandler(m *tb.Message) {
	what := &tb.Photo{File: tb.FromReader(bytes.NewReader(patheticImg))}

	if m.ReplyTo == nil {
		_, _ = gb.Send(m.Chat, what)
		return
	}

	_, _ = gb.Reply(m.ReplyTo, what)
}
