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
		"binancecoin": "BNB",
		"reserve-rights-token": "RSR",
		"stellar": "XLM",
	}

	BtnBTC = CoinSelector.Data(CoinIndexMap["bitcoin"], CoinIndexMap["bitcoin"], "bitcoin")
	BtnETH = CoinSelector.Data(CoinIndexMap["ethereum"], CoinIndexMap["ethereum"], "ethereum")
	BtnADA = CoinSelector.Data(CoinIndexMap["cardano"], CoinIndexMap["cardano"], "cardano")
	BtnXRP = CoinSelector.Data(CoinIndexMap["ripple"], CoinIndexMap["ripple"], "ripple")
	BtnUNI = CoinSelector.Data(CoinIndexMap["uniswap"], CoinIndexMap["uniswap"], "uniswap")
	BtnSUSHI = CoinSelector.Data(CoinIndexMap["sushi"], CoinIndexMap["sushi"], "sushi")
	BtnDAI = CoinSelector.Data(CoinIndexMap["dai"], CoinIndexMap["dai"], "dai")
	BtnDOGE = CoinSelector.Data(CoinIndexMap["dogecoin"], CoinIndexMap["dogecoin"], "dogecoin")
	BtnBNB = CoinSelector.Data(CoinIndexMap["binancecoin"], CoinIndexMap["binancecoin"], "binancecoin")
	BtnRSR = CoinSelector.Data(CoinIndexMap["reserve-rights-token"], CoinIndexMap["reserve-rights-token"], "reserve-rights-token")
	BtnXLM = CoinSelector.Data(CoinIndexMap["stellar"], CoinIndexMap["stellar"], "stellar")

	CoinButtonMap = map[string]*tb.Btn {
		"bitcoin": &BtnBTC,
		"ethereum": &BtnETH,
		"cardano": &BtnADA,
		"ripple": &BtnXRP,
		"uniswap": &BtnUNI,
		"sushi": &BtnSUSHI,
		"dai": &BtnDAI,
		"dogecoin": &BtnDOGE,
		"binancecoin": &BtnBNB,
		"reserve-rights-token": &BtnRSR,
		"stellar": &BtnXLM,
	}

)

func OrderCoinSelector() {
	CoinSelector.Inline(
		CoinSelector.Row(BtnBTC, BtnETH, BtnADA, BtnXRP),
		CoinSelector.Row(BtnUNI, BtnSUSHI, BtnDAI, BtnDOGE),
		CoinSelector.Row(BtnBNB, BtnRSR, BtnXLM),
	)
}