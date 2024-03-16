package usecases

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"card-validator/src/entities"
)

const (
	ERR_CODE_CARD_NUMBER_NOT_VALID           = 110
	ERR_CODE_CARD_EXPERATION_MONTH_NOT_VALID = 120
	ERR_CODE_CARD_EXPERATION_YEAR_NOT_VALID  = 130
	ERR_CODE_CARD_EPIERED                    = 140
)

var (
	ErrCardNumberNotValid                  = fmt.Errorf("card number not vatid")
	ErrCardNumberNotValidCardNumberEmpty   = fmt.Errorf("card number empty")
	ErrCardNumberNotValidInvalidCharacters = fmt.Errorf("card number contains invalid characters")
	ErrCardNumberNotValidLuhnValidation    = fmt.Errorf("card number not vatid luhn validation")
	ErrCardNumberNotValidMask              = fmt.Errorf("card number mask didn`t compose to general standarts")

	ErrCardExpirationMontNotValidOutOfRange = fmt.Errorf("card experetion month out of range")
	ErrCardExpirationYearNotValidOutOfRange = fmt.Errorf("card experetion year out of range")
	ErrCardExpired                          = fmt.Errorf("card epired")
)

type CardUseCase struct {
	log zerolog.Logger
}

func NewCardUseCase(l zerolog.Logger) CardUseCase {
	return CardUseCase{
		log: l,
	}
}

func composeNotValidCardResponse(code int, err error) entities.CardValidationResponse {
	return entities.CardValidationResponse{
		Valid: false,
		Error: &entities.Error{
			Code:    code,
			Message: err.Error(),
		},
	}
}

func (c CardUseCase) ValidateCard(card entities.Card) (entities.CardValidationResponse, error) {
	var result entities.CardValidationResponse

	if err := validateCardNumber(card.CardNumber); err != nil {
		return composeNotValidCardResponse(ERR_CODE_CARD_NUMBER_NOT_VALID, err), err
	}

	if err := validateCardMonthExpirationDate(card.ExpirationMont); err != nil {
		return composeNotValidCardResponse(ERR_CODE_CARD_EXPERATION_MONTH_NOT_VALID, err), err
	}

	if err := validateCardYearExpirationDate(card.ExpirationYear); err != nil {
		return composeNotValidCardResponse(ERR_CODE_CARD_EXPERATION_YEAR_NOT_VALID, err), err
	}

	timeNow := time.Now()
	if err := validateCardExpirationDate(card.ExpirationMont, card.ExpirationYear, timeNow); err != nil {
		return composeNotValidCardResponse(ERR_CODE_CARD_EPIERED, err), err
	}

	result.Valid = true
	return result, nil
}

func validateCardNumber(cardNumber string) error {
	// Remove any spaces in the card number
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	if cardNumber == "" {
		return ErrCardNumberNotValidCardNumberEmpty
	}

	// Check if the card number contains only digits
	if _, err := strconv.Atoi(cardNumber); err != nil {
		return ErrCardNumberNotValidInvalidCharacters
	}

	//Validate card number with luhn algorithm
	if !luhnValidationAlgorithm(cardNumber) {
		return ErrCardNumberNotValidLuhnValidation
	}

	if !validateCardFormat(cardNumber) {
		return ErrCardNumberNotValidMask
	}

	return nil
}

func luhnValidationAlgorithm(cardNumber string) bool {
	total := 0
	isSecondDigit := false

	// iterate through the card number digits in reverse order
	for i := len(cardNumber) - 1; i >= 0; i-- {
		// conver the digit character to an integer
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit

		isSecondDigit = !isSecondDigit
	}

	valid := total%10 == 0

	return valid
}

//^(?:
//(4[0-9]{12}(?:[0-9]{3})?) |          # Visa
//(5[1-5][0-9]{14}) |                  # MasterCard
//(6(?:011|5[0-9]{2})[0-9]{12}) |      # Discover
//(3[47][0-9]{13}) |                   # AMEX
//(3(?:0[0-5]|[68][0-9])[0-9]{11}) |   # Diners Club
//((?:2131|1800|35[0-9]{3})[0-9]{11})  # JCB
//)$

const CARD_MASK_REGX = `^(?:(4[0-9]{12}(?:[0-9]{3})?)|(5[1-5][0-9]{14})|(6(?:011|5[0-9]{2})[0-9]{12})|(3[47][0-9]{13})|(3(?:0[0-5]|[68][0-9])[0-9]{11})|((?:2131|1800|35[0-9]{3})[0-9]{11}))$`

// Function to validate credit/debit card number format
func validateCardFormat(cardNumber string) bool {
	regex := regexp.MustCompile(CARD_MASK_REGX)
	return regex.MatchString(cardNumber)
}

func validateCardMonthExpirationDate(cardMonthExpirationDate int) error {
	if cardMonthExpirationDate <= 0 || cardMonthExpirationDate > 12 {
		return ErrCardExpirationMontNotValidOutOfRange
	}

	return nil
}

func validateCardYearExpirationDate(cardYearExpirationDate int) error {
	if cardYearExpirationDate <= 0 || cardYearExpirationDate < 1000 {
		return ErrCardExpirationYearNotValidOutOfRange
	}
	return nil
}

func validateCardExpirationDate(cardMonthExpirationDate, cardYearExpirationDate int, now time.Time) error {
	cardExpirationDate := time.Date(cardYearExpirationDate, time.Month(cardMonthExpirationDate), 1, 0, 0, 0, 0, time.Local)

	if now.After(cardExpirationDate) {
		return ErrCardExpired
	}

	return nil
}
