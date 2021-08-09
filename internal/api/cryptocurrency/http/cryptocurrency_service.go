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
	"github.com/pkg/errors"
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

	defer func() { // recovers from panic, login errors and returning
		if r := recover(); r != nil {
			log.Panic("panic ocurred")
			err = errors.Errorf("panic ocurred: %v", r)
			return
		}
	}()

	defer r.Body.Close() //Close body reading

	var cryptocurrency cryptocurrency.Cryptocurrency

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	response = api.NewResponse(cryptocurrencyId)

	if r.StatusCode != http.StatusOK {
		return
	}

	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		return
	}
	content := api.Payload{Price: cryptocurrency.MarketData.CurrentPrice.Usd, Currency: "USD"}
	response.SetPayload(&content)
	response.Partial = false

	return response, nil
}

func (s *cryptocurrencyService) GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []api.Response) {
	cryptocurrencies := make(chan api.Response) // creates channel to receive API responses

	for _, id := range cryptocurrencyIds {
		id := id
		go func() { // go routine to call API (independently execution for each id)
			cryptocurrency, err := s.GetCryptocurrencyPrice(id, c)
			if err != nil {
				return
			}
			if cryptocurrency != nil {
				cryptocurrencies <- *cryptocurrency //sends response to channel
			}
		}()
	}

	defer close(cryptocurrencies) // closes cryptocurrnecy channel

	var cryptocurrenciesSlice []api.Response //slice to receive and present responses

	for range cryptocurrencyIds { // reading channel response and appending to response slice
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
