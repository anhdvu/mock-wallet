package main

import "net/http"

// this handler is used as placeholder during development
// it will be removed once all handlers are finished
func (app *application) homeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := jsonResponse{
			"GET -> /":                     "To view endpoint information",
			"POST -> /companion":           "To receive Companion Remote API calls",
			"POST -> /remote":              "To receive Remote Messageing API calls",
			"GET -> /companion/responses":  "To view current response codes to Companion Remote API calls",
			"PUT ->  /companion/responses": "To update one or multiple response codes to Companion Remote API calls",
			"GET -> /api/logs":             "To view API logs",
		}

		err := app.respondJSON(w, http.StatusOK, payload)
		if err != nil {
			app.serverErrorResponse(w, err)
			return
		}
	})
}