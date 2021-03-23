package coins

import tb "gopkg.in/tucnak/telebot.v2"

var (
	CoinSelector = &tb.ReplyMarkup{}

	CoinIndexMap = map[string]string {
		"bitcoin": "BTC",
		"ethereum": "ETH",
		"cardano": "ADA",
		"ripple": "XRP",
		"uniswap": "UNISWAP",
		"sushi": "SUSHI",
		"dai": "DAI",
		"dogecoin": "DOGE",
	}

	BtnBTC = CoinSelector.Data(CoinIndexMap["bitcoin"], CoinIndexMap["bitcoin"], "bitcoin")
	BtnETH = CoinSelector.Data(CoinIndexMap["ethereum"], CoinIndexMap["ethereum"], "ethereum")
	BtnADA = CoinSelector.Data(CoinIndexMap["cardano"], CoinIndexMap["cardano"], "cardano")
	BtnXRP = CoinSelector.Data(CoinIndexMap["ripple"], CoinIndexMap["ripple"], "ripple")
	BtnUNI = CoinSelector.Data(CoinIndexMap["uniswap"], CoinIndexMap["uniswap"], "uniswap")
	BtnSUSHI = CoinSelector.Data(CoinIndexMap["sushi"], CoinIndexMap["sushi"], "sushi")
	BtnDAI = CoinSelector.Data(CoinIndexMap["dai"], CoinIndexMap["dai"], "dai")
	BtnDOGE = CoinSelector.Data(CoinIndexMap["dogecoin"], CoinIndexMap["dogecoin"], "dogecoin")

	CoinButtonMap = map[string]*tb.Btn {
		"bitcoin": &BtnBTC,
		"ethereum": &BtnETH,
		"cardano": &BtnADA,
		"ripple": &BtnXRP,
		"uniswap": &BtnUNI,
		"sushi": &BtnSUSHI,
		"dai": &BtnDAI,
		"dogecoin": &BtnDOGE,
	}

)

func OrderCoinSelector() {
	CoinSelector.Inline(
		CoinSelector.Row(BtnBTC, BtnETH, BtnADA),
		CoinSelector.Row(BtnXRP, BtnUNI, BtnSUSHI),
		CoinSelector.Row(BtnDAI, BtnDOGE),
	)
}