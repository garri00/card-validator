package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"card-validator/api/handlers"
	"card-validator/api/usecases"
	"card-validator/pkg/logger"
)

type Router struct {
	chi.Router
	log zerolog.Logger
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Set timeout for incoming requests
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/health-check", handlers.HealthCheck)

	//card routes
	r.Route("/card", func(r chi.Router) {
		cardUseCase := usecases.NewCardUseCase(logger.Log)
		cardHandler := handlers.NewCardHandler(cardUseCase, logger.Log)

		r.Post("/validate", cardHandler.CardValidation)
	})

	return r
}
