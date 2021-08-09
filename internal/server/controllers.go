package server

import "github.com/gin-gonic/gin"

type CryptocurrencyController interface {
	GetCryptocurrenciesPrices(c *gin.Context)
}
