package utils

import (
	"github.com/tomassirio/bitcoinTelegram/config"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

var Bot, err = tb.NewBot(tb.Settings{
	// You can also set custom API URL.
	// If field is empty it equals to "https://api.telegram.org".
	Token:  config.LoadConfig().Token,
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})
