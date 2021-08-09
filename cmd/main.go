package main

import (
	"ejercicio-golang/internal/api"
	cryptocurrencyHttp "ejercicio-golang/internal/api/cryptocurrency/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cryptocurrencyService := cryptocurrencyHttp.NewCryptocurrencyService()

	router.GET("/myapi", func(c *gin.Context) {
		cryptocurrenciesIds := []string{"bitcoin", "ethereum", "doge"}
		response := cryptocurrencyService.GetCryptocurrenciesPrices(cryptocurrenciesIds, c)

		c.JSON(api.HasPartialResponse(response), response)
	})

	router.Run()
}
