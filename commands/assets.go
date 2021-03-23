package commands

import (
	. "github.com/tomassirio/bitcoinTelegram/coins"
	. "github.com/tomassirio/bitcoinTelegram/utils"
	tb "gopkg.in/tucnak/telebot.v2"
)

var HandleAssets = func(m *tb.Message) {

	OrderCoinSelector()

	Bot.Send(m.Sender, "Here are the current Assets", CoinSelector)

	for _, button := range CoinButtonMap {
		Bot.Handle(button, func(c *tb.Callback) {
			p := GetPrice(c.Data)

			Bot.Send(m.Chat, CoinIndexMap[c.Data] + "'s Current price is: U$S "+ F64ToString(p))
		})
	}
}