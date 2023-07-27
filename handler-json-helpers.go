package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type jsonResponse map[string]any

func (app *application) respondJSON(w http.ResponseWriter, status int, data any) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}

func (app *application) respondRaw(w http.ResponseWriter, status int, data []byte) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

	return nil
}

func (app *application) respondWithJSON(w http.ResponseWriter, status int, data any) error {
	return app.respondJSONWithLogRecord(w, status, data, nil)
}

func (app *application) respondJSONWithLogRecord(w http.ResponseWriter, status int, data any, log *logRecord) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	log.Response = string(payload)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formatted JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formatted JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain a single JSON value")
	}

	return nil
}

func (app *application) getAuthorizationHeaderChecksum(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader != "" {
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) == 2 && headerParts[0] == "CS-HMAC-SHA-256" {
			credentials := strings.Split(headerParts[1], "=")
			if len(credentials) == 3 && credentials[1] == "Terminal" {
				return credentials[2], nil
			}
			return "", errors.New("authorization header credentials were incorrectly formatted")
		}
		return "", errors.New("authorization header was incorrectly formatted")
	}
	return "", errors.New("authorization header was not found")
}
