package commands

import (
	. "github.com/tomassirio/bitcoinTelegram/coins"
	. "github.com/tomassirio/bitcoinTelegram/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"reflect"
	"strings"
)

var HandleSummary = func(m *tb.Message) {
	keys := reflect.ValueOf(CoinIndexMap).MapKeys()
	strKeys := make([]string, len(keys))
	vc := []string{"usd"}

	for _, kv := range keys {
		strKeys = append(strKeys, kv.String())
	}

	sp, err := API.SimplePrice(strKeys, vc)
	if err != nil {
		log.Fatal(err)
	}

	builder := strings.Builder{}
	builder.WriteString("This is Today's Crypto Summary\n\n")

	for coin, price := range *sp {
		if coin != ""{
			builder.WriteString(CoinIndexMap[coin] + ": U$S " + F64ToString(float64(price["usd"])) +"\n")
		}
	}

	Bot.Send(m.Chat, builder.String())

}