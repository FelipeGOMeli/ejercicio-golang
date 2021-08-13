package server

import (
	"ejercicio-golang/internal/api/cryptocurrency"
	"ejercicio-golang/internal/server"

	"net/http"

	"github.com/gin-gonic/gin"
)

type cryptocurrencyController struct {
	cryptocurrencyService cryptocurrency.CryptocurrencyService
}

func NewCryptocurrencyController(cryptocurrencyService cryptocurrency.CryptocurrencyService) server.CryptocurrencyController {
	return &cryptocurrencyController{cryptocurrencyService}
}

func (cyptocontroller *cryptocurrencyController) GetCryptocurrenciesPrices(c *gin.Context) {
	cryptocurrenciesIds := []string{"bitcoin", "ethereum", "dogecoin"}
	cryptocurrencies := cyptocontroller.cryptocurrencyService.GetCryptocurrenciesPrices(cryptocurrenciesIds, c)
	c.JSON(hasPartialResponse(cryptocurrencies), cryptocurrencies)
}

func hasPartialResponse(response []cryptocurrency.CryptocurrencyResponse) int {
	for _, r := range response {
		if r.Partial {
			return http.StatusPartialContent
		}
	}
	return http.StatusOK
}
