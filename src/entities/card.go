package entities

type Card struct {
	CardNumber     string
	ExpirationMont int
	ExpirationYear int
}

type CardValidationRequest struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationMont string `json:"expirationMont"`
	ExpirationYear string `json:"expirationYear"`
}

type CardValidationResponse struct {
	Valid bool   `json:"valid"`
	Error *Error `json:"error,omitempty"`
}
