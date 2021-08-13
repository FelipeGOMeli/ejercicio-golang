package cryptocurrency

import (
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
	GetCryptocurrencyPrice(CryptocurrencyId string, c *gin.Context) (response *CryptocurrencyResponse, err error)
	GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []CryptocurrencyResponse)
}

type CryptocurrencyResponse struct {
	Id string `json:"id"`
	Payload
	Partial bool `json:"partial"`
}

type Payload struct {
	Content interface{} `json:"content,omitempty"`
}

type CryptocurrencyContent struct {
	Price    float64 `json:"price,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

func NewResponse(id string) *CryptocurrencyResponse {
	return &CryptocurrencyResponse{
		Id:      id,
		Partial: true,
	}
}

func NewPayload(price float64, currency string) *Payload {
	return &Payload{
		Content: CryptocurrencyContent{
			Price:    price,
			Currency: currency,
		},
	}
}

func (r *CryptocurrencyResponse) SetPayload(payload Payload) {
	r.Payload = payload
	r.Partial = false
}
