package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	tb "gopkg.in/tucnak/telebot.v2"
)

func evalHandler(m *tb.Message) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			_, _ = gb.Send(m.Chat, "Error during evaluation")
		}
	}()

	expression, err := govaluate.NewEvaluableExpression(m.Payload)
	result, err := expression.Evaluate(nil)

	if err != nil {
		_, _ = gb.Send(m.Chat, err.Error())
		return
	}

	_, _ = gb.Send(m.Chat, fmt.Sprintf("%v", result))
}
