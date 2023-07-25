package main

import (
	"fmt"
	"net/http"
)

func (app *application) respondJSONWithError(w http.ResponseWriter, status int, errMsg string) {
	payload := jsonResponse{
		"error": errMsg,
	}

	err := app.respondJSON(w, status, payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, err error) {
	errMsg := "the server ran into a technical issue."
	app.logger.Println(err)
	app.respondJSONWithError(w, http.StatusInternalServerError, errMsg)
}

func (app *application) badRequestResponse(w http.ResponseWriter, err error) {
	errMsg := err.Error()

	app.respondJSONWithError(w, http.StatusBadRequest, errMsg)
}

func (app *application) notFoundHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := "the requested resource could not be found"

		app.respondJSONWithError(w, http.StatusNotFound, msg)
	})
}

func (app *application) methodNotAllowedHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("the method %s is not allowed", r.Method)

		app.respondJSONWithError(w, http.StatusMethodNotAllowed, msg)
	})
}
