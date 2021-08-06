package cryptocurrency

type Cryptocurrency struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	MarketData struct {
		CurrentPrice struct {
			Usd int `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}
