package api

type Response struct {
	Id      string `json:"id"`
	Details `json:"content"`
	Partial bool `json:"partial"`
}

type PartialResponse struct {
	Id      string `json:"id"`
	Partial bool   `json:"partial"`
}
