package main

import (
	"fmt"
	"net/http"
)

func (app *application) enforcePOSTRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only POST method is allowed")
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (app *application) enforceHTTPMethod(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Allow", method)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, fmt.Sprintf("only %s method is allowed", method))
			return
		}

		h.ServeHTTP(w, r)
	})
}
