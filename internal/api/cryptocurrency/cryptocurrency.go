package cryptocurrency

import (
	"ejercicio-golang/internal/api"

	"github.com/gin-gonic/gin"
)

type Cryptocurrency struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

type CryptocurrencyService interface {
	GetCryptocurrencyPrice(CryptocurrencyId string, c *gin.Context) (response *api.Response, err error)
	GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []api.Response)
}
