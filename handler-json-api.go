package main

import (
	"net/http"
)

func (app *application) JSONRequestHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log, err := newLogRecord("json")
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		checksum, err := app.getAuthorizationHeaderChecksum(r)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}
		log.Checksum = checksum

		payload := &jsonMessage{}
		err = app.readJSON(w, r, payload)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}

		// TO-DO
	})
}
