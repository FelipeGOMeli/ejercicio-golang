package api

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

func NewPayload(price float64, currency string) *Payload {
	return &Payload{
		Price:    price,
		Currency: currency,
	}
}

func (r *Response) SetPayload(payload *Payload) {
	r.Payload = payload
}
