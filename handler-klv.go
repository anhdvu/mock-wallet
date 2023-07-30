package main

import (
	"errors"
	"net/http"
)

func (app *application) klvHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only GET method is allowed")
			return
		}

		str := r.URL.Query().Get("string")
		if str == "" {
			app.badRequestResponse(w, errors.New("no KLV was provided"))
			return
		}

		out, err := breakDownKLV(str)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}

		res := jsonResponse{
			"message": "klv breakdown",
			"result":  out,
		}

		err = app.respondJSON(w, http.StatusOK, res)
		if err != nil {
			app.serverErrorResponse(w, err)
		}
	}
}
