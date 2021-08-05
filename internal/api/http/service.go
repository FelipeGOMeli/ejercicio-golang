package http

import (
	"ejercicio-golang/internal/api"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResponse() {
	router := gin.Default()
	externalUrl := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
	url := "/myapi"

	router.GET(url, func(c *gin.Context) {

		response, err := http.Get(externalUrl)

		if err != nil {
			log.Fatalln("Failed to call service")
		}

		defer response.Body.Close()

		var bitcoinResponse api.BitcoinResponse

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatalln("Failed to read response")
		}

		if response.StatusCode != http.StatusOK {
			parsedResponse := api.PartialResponse{Id: "btc", Partial: false}
			c.JSON(http.StatusPartialContent, parsedResponse)
		} else {
			_ = json.Unmarshal(body, &bitcoinResponse)
			content := api.Details{Price: float64(bitcoinResponse.Bitcoin.Usd), Currency: "usd"}
			parsedResponse := api.Response{Id: "btc", Details: content, Partial: true}
			c.JSON(http.StatusOK, parsedResponse)
		}

	})
	router.Run()
}
