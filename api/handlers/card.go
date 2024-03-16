package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"card-validator/src/entities"
)

type CardValidationUseCase interface {
	ValidateCard(card entities.Card) (entities.CardValidationResponse, error)
}

type CarValidationHandler struct {
	cardValidationUseCase CardValidationUseCase
	log                   zerolog.Logger
}

func NewCardHandler(cardValidationUseCase CardValidationUseCase, l zerolog.Logger) CarValidationHandler {
	return CarValidationHandler{
		cardValidationUseCase: cardValidationUseCase,
		log:                   l,
	}
}

func (c CarValidationHandler) CardValidation(w http.ResponseWriter, r *http.Request) {
	var cardRequest entities.CardValidationRequest

	if err := decodeBody(r.Body, &cardRequest); err != nil {
		respondErr(w, c.log, err, http.StatusInternalServerError)
		return
	}

	card, err := cardRequestToCard(cardRequest)
	if err != nil {
		respondErr(w, c.log, err, http.StatusBadRequest)
		return
	}

	resp, err := c.cardValidationUseCase.ValidateCard(card)
	if err != nil {
		respond(w, c.log, resp)
		return
	}

	respond(w, c.log, resp)

}

func cardRequestToCard(cardRequest entities.CardValidationRequest) (entities.Card, error) {
	var card entities.Card

	card.CardNumber = strings.TrimSpace(cardRequest.CardNumber)

	expirationMonth, err := strconv.Atoi(cardRequest.ExpirationMont)
	if err != nil {
		return card, fmt.Errorf("convert expirationMonth failed: %w", err)
	}

	card.ExpirationMont = expirationMonth

	expirationYear, err := strconv.Atoi(cardRequest.ExpirationYear)
	if err != nil {
		return card, fmt.Errorf("convert expirationYear failed: %w", err)
	}

	card.ExpirationYear = expirationYear

	return card, nil
}
