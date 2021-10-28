package main

import (
	"log"
	"os"
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

var owmApiKey = func() (key string) {
	key = os.Getenv("OWM_API_KEY")

	if args := os.Args[2:]; len(args) != 0 {
		key = args[0]
	}

	if key == "" {
		log.Fatal("provide telegram bot api key as first argument")
	}

	return key
}()
