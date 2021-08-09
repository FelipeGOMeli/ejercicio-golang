package http

import (
	"ejercicio-golang/internal/api"
	"ejercicio-golang/internal/api/cryptocurrency"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (s *cryptocurrencyService) GetCryptocurrencyPrice(cryptocurrencyId string, c *gin.Context) (response *api.Response, err error) {

	externalUrl := setPath(cryptocurrencyId)
	r, err := http.Get(externalUrl)

	if err != nil {
		return
	}

	defer r.Body.Close() //Close body reading

	var cryptocurrency cryptocurrency.Cryptocurrency

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	parsedResponse := api.Response{}

	if r.StatusCode != http.StatusOK {
		parsedResponse = api.Response{Id: cryptocurrencyId, Partial: true}
		return &parsedResponse, nil
	}

	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		return
	}
	content := api.Payload{Price: cryptocurrency.MarketData.CurrentPrice.Usd, Currency: "USD"}
	parsedResponse = api.Response{Id: cryptocurrency.Symbol, Payload: &content, Partial: false}

	return &parsedResponse, nil
}

func (s *cryptocurrencyService) GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []api.Response) {
	cryptocurrencies := make(chan api.Response) // create channel to receive API responses

	for _, id := range cryptocurrencyIds {
		id := id
		go func() { // go routine to call API (independently execution for each id)
			cryptocurrency, err := s.GetCryptocurrencyPrice(id, c)
			if err != nil {
				return
			}
			if cryptocurrency != nil {
				cryptocurrencies <- *cryptocurrency //sent response to channel
			}
		}()
	}

	var cryptocurrenciesSlice []api.Response //slice to receive and present responses
	timeout := time.After(1 * time.Second)

	for range cryptocurrencyIds { // reading channel response and appending to response slice
		select {
		case result := <-cryptocurrencies:
			cryptocurrenciesSlice = append(cryptocurrenciesSlice, result)
		case <-timeout:
			fmt.Println("timed out")
		}
	}
	return cryptocurrenciesSlice
}

func setPath(cryptocurrencyId string) string {
	externalUrl := url.URL{Scheme: "https", Host: "api.coingecko.com"}
	externalUrl.Path = fmt.Sprintf("/api/v3/coins/%s", cryptocurrencyId)
	query := externalUrl.Query()
	query.Set("localization", "false")
	query.Set("tickers", "false")
	query.Set("community_data", "false")
	query.Set("developer_data", "false")
	query.Set("sparkline", "false")
	externalUrl.RawQuery = query.Encode()
	return externalUrl.String()
}
