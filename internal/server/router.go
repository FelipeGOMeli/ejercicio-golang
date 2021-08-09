package server

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(cryptocurrencyController CryptocurrencyController) *Router {
	router := gin.Default()

	router.GET("/myapi", func(c *gin.Context) {
		cryptocurrencyController.GetCryptocurrenciesPrices(c)
	})

	return &Router{
		Engine: router,
	}
}

func (r *Router) Run(serverAdress string) error {
	return r.Engine.Run(serverAdress)
}
