package http

import (
	"ejercicio-golang/internal/api/cryptocurrency"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const currency = "USD"

type cryptocurrencyService struct {
	endpoint url.URL
}

func NewCryptocurrencyService(endpoint url.URL) cryptocurrency.CryptocurrencyService {
	return &cryptocurrencyService{endpoint: endpoint}
}

func (s *cryptocurrencyService) GetCryptocurrencyPrice(cryptocurrencyId string, c *gin.Context) (response *cryptocurrency.CryptocurrencyResponse, err error) {
	response = cryptocurrency.NewResponse(cryptocurrencyId)
	externalEndpoint := setCryptocurrencyId(cryptocurrencyId, s.endpoint)
	r, err := http.Get(externalEndpoint)

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

	var cryptocurrencyPayload cryptocurrency.Cryptocurrency

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	if r.StatusCode != http.StatusOK {
		return
	}

	err = json.Unmarshal(body, &cryptocurrencyPayload)
	if err != nil {
		panic("failed to parse json")
	}
	content := cryptocurrency.NewPayload(cryptocurrencyPayload.MarketData.CurrentPrice.Usd, currency)
	response.SetPayload(*content)

	return response, nil
}

func (s *cryptocurrencyService) GetCryptocurrenciesPrices(cryptocurrencyIds []string, c *gin.Context) (response []cryptocurrency.CryptocurrencyResponse) {
	cryptocurrencies := make(chan cryptocurrency.CryptocurrencyResponse, len(cryptocurrencyIds)) //creates channel to receive API responses

	var wg sync.WaitGroup
	wg.Add(len(cryptocurrencyIds))

	for _, id := range cryptocurrencyIds {
		id := id
		go func() { //go routine to call API (independently execution for each id)
			defer wg.Done() //after execution decrements wg counter
			cryptocurrencyResponse, err := s.GetCryptocurrencyPrice(id, c)
			if err != nil {
				fmt.Printf("failed to get %v price\n", id)
			}
			if cryptocurrencyResponse != nil {
				cryptocurrencies <- *cryptocurrencyResponse //sends response to channel
			}
		}()
	}

	wg.Wait()               //waits until all requests finishes
	close(cryptocurrencies) //closes cryptocurrnecy channel

	var cryptocurrenciesSlice []cryptocurrency.CryptocurrencyResponse //slice to receive and show responses

	for range cryptocurrencyIds { //reading channel response and appending to responses slice
		result := <-cryptocurrencies
		cryptocurrenciesSlice = append(cryptocurrenciesSlice, result)
	}
	return cryptocurrenciesSlice
}

func setCryptocurrencyId(cryptocurrencyId string, endpoint url.URL) string {
	endpoint.Path = path.Join(endpoint.Path, cryptocurrencyId)

	return endpoint.String()
}
