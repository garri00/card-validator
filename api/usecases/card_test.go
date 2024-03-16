package usecases

import (
	"os"
	"testing"

	"github.com/rs/zerolog"

	"card-validator/src/entities"
)

const (
	VALID_MONT = 1
	VALID_YEAR = 2030
)

//
//American Express  371449635398431
//Diners Club  30569309025904
//Discover  6011111111111117
//JCB  3530111333300000
//MasterCard  5555555555554444
//Visa  4111111111111111

var validCards = []entities.Card{
	{
		CardNumber:     "4111111111111111",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		CardNumber:     "6011000990139424",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		CardNumber:     "3530111333300000",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		CardNumber:     "6011111111111117",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		CardNumber:     "30569309025904",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		CardNumber:     "371449635398431",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
}

var notValidCards = []entities.Card{
	{
		//Invalid card num
		CardNumber:     "5610591081018260",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		//Invalid card num
		CardNumber:     "5610591081018250",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		//Invalid card num
		CardNumber:     "12345678901234567",
		ExpirationMont: VALID_MONT,
		ExpirationYear: VALID_YEAR,
	},
	{
		//Invalid expiration month
		CardNumber:     "4111111111111111",
		ExpirationMont: 00,
		ExpirationYear: VALID_YEAR,
	},
	{
		//Invalid expiration year
		CardNumber:     "4111111111111111",
		ExpirationMont: VALID_MONT,
		ExpirationYear: 999,
	},
	{
		//Card expierd
		CardNumber:     "4111111111111111",
		ExpirationMont: 01,
		ExpirationYear: 2020,
	},
}

func TestCardUseCase_ValidateCard(t *testing.T) {

	Log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	tests := []struct {
		name  string
		log   zerolog.Logger
		cards []entities.Card
		//want    entities.CardValidationResponse
		wantErr bool
	}{
		{
			name:  "ValidCards",
			log:   Log,
			cards: validCards,
			//want:    entities.CardValidationResponse{},
			wantErr: false,
		},
		{
			name:  "ValidCards",
			log:   Log,
			cards: notValidCards,
			//want:    entities.CardValidationResponse{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CardUseCase{
				log: Log,
			}

			for _, card := range tt.cards {
				_, err := c.ValidateCard(card)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateCard() error = %v, wantErr: %v, card:%s", err, tt.wantErr, card.CardNumber)
				}
			}

		})
	}
}
