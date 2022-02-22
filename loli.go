package main

import (
	"encoding/json"
	tb "gopkg.in/tucnak/telebot.v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AtfItem struct {
	FileUrl string `json:"file_url"`
}

func loliHandler(m *tb.Message) {
	req, err := getLoliReq(m.Payload)
	if err != nil {
		somethingWentWrong(m)
		return
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		somethingWentWrong(m)
		return
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		somethingWentWrong(m)
		return
	}

	var items []AtfItem

	if jsonErr := json.Unmarshal(body, &items); jsonErr != nil || len(items) == 0 {
		_, _ = gb.Send(m.Chat, "No lolis found 😔")
		return
	}

	_, _ = gb.Send(m.Chat, &tb.Photo{File: tb.FromURL(items[0].FileUrl)})
}

func getLoliReq(payload string) (*http.Request, error) {
	tags := "rating:safe order:random loli 1girl -toddler"

	if payload != "" {
		tags += " *" + strings.ReplaceAll(payload, " ", "_") + "*"
	}

	queryParams := url.Values{
		"tags":  {tags},
		"limit": {"1"},
	}

	return http.NewRequest(
		http.MethodGet,
		"https://booru.allthefallen.moe/posts.json?"+queryParams.Encode(),
		nil,
	)
}
