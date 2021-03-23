package main

import (
	. "github.com/tomassirio/bitcoinTelegram/handler"
	. "github.com/tomassirio/bitcoinTelegram/utils"
	"log"
)

func main() {

	for k, v := range LoadHandler() {
		Bot.Handle(k, v)
		log.Println(k + " ✅  Loaded!")
	}

	Bot.Start()

}
