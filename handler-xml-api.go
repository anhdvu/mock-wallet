package main

import (
	"context"
	"net/http"
	"time"
)

func (app *application) XMLRequestHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log, err := newLogRecord("xml")
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		payload := &xmlPayload{}
		err = app.readXML(r, payload, log)
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

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = app.apiLogger.SaveLog(ctx, log)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		err = app.respondXML(w, http.StatusOK, response)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
		}
	})
}
