package commands

import (
	. "github.com/tomassirio/bitcoinTelegram/utils"
	"log"
)

func GetPrice(coin string) float64 {
	p, err := API.SimpleSinglePrice(coin, "usd")
	if err != nil {
		log.Println("The desired coin is not listed on the API list")
	}
	return float64(p.MarketPrice)
}
