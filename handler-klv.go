package main

import (
	"net/http"
)

func (app *application) klvHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only POST method is allowed")
			return
		}

		var input struct {
			KLV string `json:"klv"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			app.badRequestResponse(w, err)
			return
		}

		out, err := breakDownKLV(input.KLV)
		if err != nil {
			app.logger.Println(err)
			app.badRequestResponse(w, err)
			return
		}

		res := jsonResponse{
			"klv":    input.KLV,
			"result": out,
		}

		err = app.respondJSON(w, http.StatusOK, res)
		if err != nil {
			app.serverErrorResponse(w, err)
		}
	}
}
