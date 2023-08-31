package main

import (
	"context"
	"net/http"
	"time"
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

		p := &jsonMessage{}
		err = app.readJSONWithLogRecord(w, r, p, log)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}

		bd, err := breakDownKLV(p.Challenge)
		if err != nil {
			app.logger.Println(err)
			log.KLVBreakdown = nil
		} else {
			log.KLVBreakdown = bd
		}

		p.ResultCode = "0000"

		err = app.respondJSONWithLogRecord(w, http.StatusOK, p, log)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = app.apiLogger.SaveLog(ctx, log)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}
	})
}
