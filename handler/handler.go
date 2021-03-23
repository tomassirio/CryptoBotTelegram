package handler

import (
	. "github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)


func LoadHandler() map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/historic"] = HandleHistoric
	commandMap["/summary"] = HandleSummary
	commandMap["/assets"] = HandleAssets

	return commandMap
}
