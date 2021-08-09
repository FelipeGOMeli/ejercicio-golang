package api

import "net/http"

type Response struct {
	Id       string `json:"id"`
	*Payload `json:"content,omitempty"`
	Partial  bool `json:"partial"`
}

type Payload struct {
	Price    float64 `json:"price,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

func NewResponse(id string) *Response {
	return &Response{
		Id:      id,
		Partial: true,
	}
}

func (r *Response) SetPayload(payload *Payload) {
	r.Payload = payload
}

func HasPartialResponse(response []Response) int {
	for _, r := range response {
		if r.Partial {
			return http.StatusPartialContent
		}
	}
	return http.StatusOK
}
