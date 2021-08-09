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
	"time"

	"github.com/gin-gonic/gin"
)

type cryptocurrencyService struct {
}

func NewCryptocurrencyService() cryptocurrency.CryptocurrencyService {
	return &cryptocurrencyService{}
}

func (s *cryptocurrencyService) GetCryptocurrencyPrice(cryptocurrencyId string, c *gin.Context) (response *api.Response) {

	externalUrl := url.URL{Scheme: "https", Host: "api.coingecko.com"}
	externalUrl.Path = fmt.Sprintf("/api/v3/coins/%s", cryptocurrencyId)
	query := externalUrl.Query()
	query.Set("localization", "false")
	query.Set("tickers", "false")
	query.Set("community_data", "false")
	query.Set("developer_data", "false")
	query.Set("sparkline", "false")
	externalUrl.RawQuery = query.Encode()

	r, err := http.Get(externalUrl.String())

	if err != nil {
		log.Fatalln("Failed to call service")
	}

	defer r.Body.Close()

	var cryptocurrency cryptocurrency.Cryptocurrency

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("Failed to read response")
	}

	parsedResponse := api.Response{}

	if r.StatusCode != http.StatusOK {
		parsedResponse = api.Response{Id: cryptocurrency.Symbol, Partial: true}
	} else {
		_ = json.Unmarshal(body, &cryptocurrency)
		content := api.Details{Price: cryptocurrency.MarketData.CurrentPrice.Usd, Currency: "usd"}
		parsedResponse = api.Response{Id: cryptocurrency.Symbol, Details: content, Partial: false}
	}
	return &parsedResponse
}

func (s *cryptocurrencyService) GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) []api.Response {
	cryptocurrencies := make(chan api.Response) // create channel to receive API responses

	for _, id := range cryptocurrencyIds {
		id := id
		go func() { // go routine to call API (independently execution for each id)
			cryptocurrencies <- *s.GetCryptocurrencyPrice(id, c)
		}()
	}

	var cryptocurrenciesSlice []api.Response //slice to receive and present responses
	timeout := time.After(1 * time.Second)

	for i := 0; i < 3; i++ { // reading channel response and appending to response slice
		select {
		case result := <-cryptocurrencies:
			cryptocurrenciesSlice = append(cryptocurrenciesSlice, result)
		case <-timeout:
			fmt.Println("timed out")
		}
	}
	return cryptocurrenciesSlice
}
