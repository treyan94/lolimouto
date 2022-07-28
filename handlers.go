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
	msg := strings.ToUpper(m.Payload)

	if msg == "" {
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

	replyText = strings.TrimPrefix(replyText, "Did you mean: \n")
	mustCompile := regexp.MustCompile(split[1])
	replyMessage := "`Did you mean:` \n" + mustCompile.ReplaceAllString(replyText, strings.Join(split[2:], "/"))

	_, _ = gb.Reply(replyTo, replyMessage, &tb.SendOptions{ParseMode: tb.ModeMarkdown})
}

//go:embed pathetic.png
var patheticImg []byte

func patheticHandler(m *tb.Message) {
	mediaRes(m, &tb.Photo{File: tb.FromReader(bytes.NewReader(patheticImg))})
}

//go:embed blasphemy.mp4
var blasphemyVid []byte

func blasphemyHandler(m *tb.Message) {
	mediaRes(m, &tb.Video{File: tb.FromReader(bytes.NewReader(blasphemyVid))})
}

//go:embed "hot.mp4"
var hotVid []byte

func hotHandler(m *tb.Message) {
	mediaRes(m, &tb.Video{File: tb.FromReader(bytes.NewReader(hotVid))})
}

func mediaRes(m *tb.Message, what interface{}) {
	if m.ReplyTo == nil {
		_, _ = gb.Send(m.Chat, what)
		return
	}

	_, _ = gb.Reply(m.ReplyTo, what)

}

func weatherHandler(m *tb.Message) {
	c := weatherClient.Search(m.Payload)

	_, _ = gb.Send(m.Chat, c.Text(), &tb.SendOptions{ParseMode: tb.ModeMarkdown})
}

func rollHandler(m *tb.Message) {
	strSplit := strings.Split(strings.ToLower(m.Payload), "d")

	dice, sides, err := parseRoll(strSplit)

	if err != nil {
		_, _ = gb.Send(m.Chat, "Invalid roll")
		return
	}

	var rolls []int

	for i := 0; i < dice; i++ {
		roll := rand.Intn(sides) + 1
		rolls = append(rolls, roll)
	}

	_, _ = gb.Send(m.Chat, implodeIntSlice(rolls, ", "))
}

func parseRoll(strSplit []string) (dice int, sides int, err error) {
	if len(strSplit) == 1 {
		dice = 1
		sides, err = strconv.Atoi(strSplit[0])
		return
	}

	dice, err = strconv.Atoi(strSplit[0])
	if err != nil {
		return
	}

	sides, err = strconv.Atoi(strSplit[1])
	return
}
