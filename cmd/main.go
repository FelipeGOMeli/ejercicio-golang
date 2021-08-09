package main

import (
	cryptocurrencyHttp "ejercicio-golang/internal/api/cryptocurrency/http"
	"ejercicio-golang/internal/server"
	cryptoController "ejercicio-golang/internal/server/cryptocurrency"
)

const serverAddress = ":8080"

func main() {
	cryptocurrencyService := cryptocurrencyHttp.NewCryptocurrencyService()
	cryptocurrencyController := cryptoController.NewCryptocurrencyController(cryptocurrencyService)

	router := server.NewRouter(cryptocurrencyController)

	router.Run(serverAddress)
}
