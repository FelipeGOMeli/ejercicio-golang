package api

type BitcoinResponse struct {
	Bitcoin struct {
		Usd int `json:"usd"`
	} `json:"bitcoin"`
}
