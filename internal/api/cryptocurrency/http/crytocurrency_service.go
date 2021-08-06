package http

import (
	"ejercicio-golang/internal/api"
	"ejercicio-golang/internal/api/cryptocurrency"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetCryptocurrencyPrice(cryptocurrencyId string) {
	router := gin.Default()
	externalUrl := url.URL{Scheme: "https", Host: "api.coingecko.com"}
	externalUrl.Path = fmt.Sprintf("/api/v3/coins/%s", cryptocurrencyId)
	query := externalUrl.Query()
	query.Set("localization", "false")
	query.Set("tickers", "false")
	query.Set("community_data", "false")
	query.Set("developer_data", "false")
	query.Set("sparkline", "false")
	externalUrl.RawQuery = query.Encode()
	fmt.Println(externalUrl.String())
	url := "/myapi"

	router.GET(url, func(c *gin.Context) {

		response, err := http.Get(externalUrl.String())

		if err != nil {
			log.Fatalln("Failed to call service")
		}

		defer response.Body.Close()

		var cryptocurrency cryptocurrency.Cryptocurrency

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatalln("Failed to read response")
		}

		if response.StatusCode != http.StatusOK {
			parsedResponse := api.PartialResponse{Id: cryptocurrency.Symbol, Partial: false}
			c.JSON(http.StatusPartialContent, parsedResponse)
		} else {
			_ = json.Unmarshal(body, &cryptocurrency)
			fmt.Println(float64(cryptocurrency.MarketData.CurrentPrice.Usd))
			content := api.Details{Price: float64(cryptocurrency.MarketData.CurrentPrice.Usd), Currency: "usd"}
			parsedResponse := api.Response{Id: cryptocurrency.Symbol, Details: content, Partial: true}
			c.JSON(http.StatusOK, parsedResponse)
		}

	})
	router.Run()
}
