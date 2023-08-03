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

		err = app.respondJSONWithLogRecord(w, 200, p, log)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
		}
	})
}
