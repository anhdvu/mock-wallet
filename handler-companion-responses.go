package main

import "net/http"

func (app *application) companionResponsesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := app.respondJSON(w, http.StatusOK, app.companion.allResponses())
			if err != nil {
				app.logger.Println(err)
				app.serverErrorResponse(w, err)
				return
			}

		case http.MethodPut:
			input := make(map[string]int)

			err := app.readJSON(w, r, &input)
			if err != nil {
				app.logger.Println(err)
				app.badRequestResponse(w, err)
				return
			}

			for k, v := range input {
				err := app.companion.updateResponseCode(k, v)
				if err != nil {
					app.logger.Println(err)
					app.badRequestResponse(w, err)
					return
				}
			}

			err = app.respondJSON(w, http.StatusAccepted, app.companion.allResponses())
			if err != nil {
				app.logger.Println(err)
				app.serverErrorResponse(w, err)
				return
			}

		default:
			w.Header()["Allow"] = []string{http.MethodGet, http.MethodPut}
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only GET and PUT method are allowed")
			return
		}
	}
}
