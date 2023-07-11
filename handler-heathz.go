package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthz() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, fmt.Sprintf("only method %s is allowed", http.MethodGet))
			return
		}

		payload := jsonResponse{
			"status": "service is running",
		}

		err := app.respondJSON(w, 200, payload)
		if err != nil {
			app.serverErrorResponse(w, err)
			return
		}
	})
}
