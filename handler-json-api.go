package main

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid/v5"
)

func (app *application) JSONRequestHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewV4()
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		log := &record{
			Type: "json",
			Time: time.Now(),
			ID:   id,
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
	})
}
