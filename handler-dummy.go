package main

import "net/http"

// this handler is used as placeholder during development
// it will be removed once all handlers are finished
func (app *application) dummyHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := jsonResponse{
			"message": "service is under construction",
		}

		err := app.respondJSON(w, http.StatusOK, payload)
		if err != nil {
			app.serverErrorResponse(w, err)
			return
		}
	})
}
