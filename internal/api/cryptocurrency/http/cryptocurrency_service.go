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
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const currency = "USD"

type cryptocurrencyService struct {
}

func NewCryptocurrencyService() cryptocurrency.CryptocurrencyService {
	return &cryptocurrencyService{}
}

func (s *cryptocurrencyService) GetCryptocurrencyPrice(cryptocurrencyId string, c *gin.Context) (response *api.Response, err error) {
	response = api.NewResponse(cryptocurrencyId)
	externalUrl := setPath(cryptocurrencyId)
	r, err := http.Get(externalUrl)

	if err != nil {
		return
	}

	defer func() { //recovers from panic (any error that stops program flow), login errors and returning
		if r := recover(); r != nil {
			log.Panic("panic ocurred")
			err = errors.Errorf("panic ocurred: %v", r)
			return
		}
	}()

	defer r.Body.Close() //Closes body reading after

	var cryptocurrency cryptocurrency.Cryptocurrency

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	if r.StatusCode != http.StatusOK {
		return
	}

	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		panic("failed to parse file")
	}
	content := api.NewPayload(cryptocurrency.MarketData.CurrentPrice.Usd, currency)
	response.SetPayload(content)
	response.Partial = false

	return response, nil
}

func (s *cryptocurrencyService) GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []api.Response) {
	cryptocurrencies := make(chan api.Response, len(cryptocurrencyIds)) //creates channel to receive API responses

	var wg sync.WaitGroup
	wg.Add(len(cryptocurrencyIds))

	for _, id := range cryptocurrencyIds {
		id := id
		go func() { //go routine to call API (independently execution for each id)
			defer wg.Done() //after execution decrements wg counter
			cryptocurrency, err := s.GetCryptocurrencyPrice(id, c)
			if err != nil {
				fmt.Printf("failed to get %v price\n", id)
			}
			if cryptocurrency != nil {
				cryptocurrencies <- *cryptocurrency //sends response to channel
			}
		}()
	}

	wg.Wait()               //waits until all requests finishes
	close(cryptocurrencies) //closes cryptocurrnecy channel

	var cryptocurrenciesSlice []api.Response //slice to receive and show responses

	for range cryptocurrencyIds { //reading channel response and appending to responses slice
		result := <-cryptocurrencies
		cryptocurrenciesSlice = append(cryptocurrenciesSlice, result)
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
