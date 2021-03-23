package commands

import (
	. "github.com/tomassirio/bitcoinTelegram/coins"
	. "github.com/tomassirio/bitcoinTelegram/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var (
	selector = &tb.ReplyMarkup{}
	btnYes = selector.Data("Yesterday", "Yes", "yesterday")
	btnWee = selector.Data("Week", "Wee", "last week")
	btnMon = selector.Data("Month", "Mon", "last month")
	btnYea = selector.Data("Year", "Yea", "last year")
	yes = time.Now().AddDate(0, 0, -1).Format("02-01-2006")
	wee = time.Now().AddDate(0, 0, -7).Format("02-01-2006")
	mon = time.Now().AddDate(0, -1, 0).Format("02-01-2006")
	yea = time.Now().AddDate(-1, 0, 0).Format("02-01-2006")
	dateButtonMap = map[string]*tb.Btn {"yesterday": &btnYes, "week": &btnWee, "month": &btnMon, "year": &btnYea}
	dateMap = map[string]string {"yesterday": yes, "week": wee, "month": mon, "year": yea}
)

var HandleHistoric = func(m *tb.Message) {

	OrderCoinSelector()

	Bot.Send(m.Chat, "For which asset do you want to know the historic value?", CoinSelector)

	for _, button := range CoinButtonMap {
		Bot.Handle(button, func(c *tb.Callback) {
			p := GetPrice(c.Data)

			replyDate(c.Data, p, m)
		})
	}
}

func getHistoric(coin string, date string) float64{
	p, err := API.CoinsIDHistory(coin, date, false)
	if err != nil {
		log.Println("There's no data for this coin/date")
	}
	return p.MarketData.CurrentPrice["usd"]
}

func replyDate(coin string, p float64, m *tb.Message) {
	selector.Inline(
		selector.Row(btnYes, btnWee),
		selector.Row(btnMon, btnYea),
	)

	Bot.Send(m.Chat, "What's the historic you want to get?", selector)

	for _, button := range dateButtonMap {
		Bot.Handle(button, func(c *tb.Callback) {

			h := getHistoric(coin, dateMap[c.Data])

			perc := ((p-h) / h) * 100

			Bot.Send(m.Chat, CoinIndexMap[coin]+"'s " + c.Data + "'s price was: U$S "+ F64ToString(h) +
				"\nThat's a " + returnEmoji(perc) + " " + F64ToString(perc) + "% compared to Today's U$S " + F64ToString(p))
		})
	}
}

func returnEmoji(perc float64) string {
	if perc == 0.0000 {
		return "➡️"
	}
	if perc < 0.0000 {
		return "↘️"
	}
	return "↗️"
}