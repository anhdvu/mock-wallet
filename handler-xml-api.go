package main

import "net/http"

func (app *application) handleXMLRequests() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
