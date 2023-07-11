package main

import (
	"errors"
	"mime"
	"net/http"
)

func (app *application) enforceJSONHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			app.badRequestResponse(w, errors.New("Content-Type was not set"))
			return
		}

		if contentType != "" {
			mediaType, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				app.badRequestResponse(w, err)
				return
			}

			if mediaType != "application/json" {
				app.respondJSONWithError(w, http.StatusUnsupportedMediaType, "only application/json is supported")
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}
