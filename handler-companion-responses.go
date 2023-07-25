package main

import "net/http"

func (app *application) showCompanionResponsesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := app.respondJSON(w, 200, app.companion.allResponses())
			if err != nil {
				app.logger.Println(err)
				app.serverErrorResponse(w, err)
				return
			}

		case http.MethodPut:
			input := make(map[string]any)

			err := app.readJSON(w, r, &input)
			if err != nil {
				app.logger.Println(err)
				app.badRequestResponse(w, err)
				return
			}

			for k, v := range input {
				code := v.(int)

				err := app.companion.updateResponseCode(k, code)
				if err != nil {
					app.logger.Println(err)
					app.badRequestResponse(w, err)
					return
				}
			}
		default:
			w.Header()["Allow"] = []string{http.MethodGet, http.MethodPut}
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only GET and PUT method are allowed")
			return
		}
	}
}
