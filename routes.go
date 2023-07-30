package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/companion/responses", app.companionResponsesHandler())

	mux.Handle("/companion", app.enforcePOSTRequest(app.enforceXMLHeader(app.XMLRequestHandler())))
	mux.Handle("/remote", app.enforcePOSTRequest(app.enforceJSONHeader(app.JSONRequestHandler())))

	mux.Handle("/api/health", app.healthz())

	mux.Handle("/api/logs", app.viewLogsHandler())

	mux.Handle("/", app.homeHandler())

	return mux
}
