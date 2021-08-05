package http

import (
	"ejercicio-golang/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResponse() {
	router := gin.Default()
	url := "/myapi"

	router.GET(url, func(c *gin.Context) {
		queryParam := c.Query("data") //declare var queryParam and initializes it with received query param
		response := api.Response{Data: queryParam}

		c.JSON(http.StatusOK, response)
	})
	router.Run()
}
