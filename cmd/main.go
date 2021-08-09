package main

import (
	cryptocurrencyHttp "ejercicio-golang/internal/api/cryptocurrency/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cryptocurrencyService := cryptocurrencyHttp.NewCryptocurrencyService()

	router.GET("/myapi", func(c *gin.Context) {
		cryptocurrenciesIds := []string{"bitcoin", "ethereum", "dogecoin"}
		response := cryptocurrencyService.GetCryptocurrenciesPrices(cryptocurrenciesIds, c)
		c.JSON(http.StatusOK, response)
	})

	router.Run()
}
