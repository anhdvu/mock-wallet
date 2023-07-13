package main

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid/v5"
)

func (app *application) XMLRequestHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewV4()
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		log := &record{
			Type: "xml",
			Time: time.Now(),
			ID:   id,
		}

		payload := &xmlPayload{}
		err = app.readXML(r, payload)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}

		err = app.processXMLPayload(payload, log)
		if err != nil {
			app.logger.Println(err)
		}

		response, err := app.companion.makeResponse(payload.MethodName)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}
		log.Response = response
		err = app.respondXML(w, http.StatusOK, response)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
		}
	})
}
