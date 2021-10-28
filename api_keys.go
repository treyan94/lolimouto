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
		log.Fatal("Set the 'LOLIMOUTO_BOT_KEY' env variable or provide telegram bot api key as the first argument")
	}

	return key
}()

var owmApiKey = func() (key string) {
	key = os.Getenv("OWM_API_KEY")

	if key != "" {
		return key
	}

	if args := os.Args[2:]; len(args) != 0 {
		key = args[0]
	}

	if key == "" {
		log.Fatal("Set the 'OWM_API_KEY' env variable or provide telegram bot api key as the second argument")
	}

	return key
}()
