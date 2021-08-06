package main

import "ejercicio-golang/internal/api/cryptocurrency/http"

func main() {
	http.GetCryptocurrencyPrice("bitcoin")
}
