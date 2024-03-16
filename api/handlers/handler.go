package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func respond(rw http.ResponseWriter, log zerolog.Logger, data any) {
	if data == nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	var buf bytes.Buffer
	if err := encodeBody(&buf, data); err != nil {
		err = fmt.Errorf("encoding to buffer: %w", err)
		respondErr(rw, log, err, http.StatusInternalServerError)
		return
	}

	if _, err := buf.WriteTo(rw); err != nil {
		err = fmt.Errorf("writing response: %w", err)
		respondErr(rw, log, err, http.StatusInternalServerError)
		return
	}
}

func respondErr(rw http.ResponseWriter, log zerolog.Logger, err error, statusCode int) {
	log.Debug().Err(err).Msg("responding with error")
	http.Error(rw, err.Error(), statusCode)
}

// decodeBody reads data from a body and converts it to any
func decodeBody(body io.Reader, data any) error {
	if err := json.NewDecoder(body).Decode(data); err != nil {
		return fmt.Errorf("decoding body: %w", err)
	}

	return nil
}

// encodeBody writes data to a writer after converting it to JSON
func encodeBody(rw io.Writer, data any) error {
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		return fmt.Errorf("encoding body: %w", err)
	}

	return nil
}
