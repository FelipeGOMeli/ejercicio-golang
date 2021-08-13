package main

import (
	cryptocurrencyHttp "ejercicio-golang/internal/api/cryptocurrency/http"
	"ejercicio-golang/internal/server"
	cryptoController "ejercicio-golang/internal/server/cryptocurrency"
	"net/url"
)

const serverAddress = ":8080"

func main() {
	externalEndpoint := configExternalEndpoint()
	cryptocurrencyService := cryptocurrencyHttp.NewCryptocurrencyService(externalEndpoint)
	cryptocurrencyController := cryptoController.NewCryptocurrencyController(cryptocurrencyService)

	router := server.NewRouter(cryptocurrencyController)

	router.Run(serverAddress)
}

func configExternalEndpoint() (externalUrl url.URL) {
	externalUrl = url.URL{Scheme: "https", Host: "api.coingecko.com"}
	externalUrl.Path = "/api/v3/coins/"
	query := externalUrl.Query()
	query.Set("localization", "false")
	query.Set("tickers", "false")
	query.Set("community_data", "false")
	query.Set("developer_data", "false")
	query.Set("sparkline", "false")

	return externalUrl
}
