package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/companion", app.enforcePOSTRequest(app.enforceXMLHeader(app.dummyHandler())))
	mux.Handle("/remote", app.enforcePOSTRequest(app.enforceJSONHeader(app.dummyHandler())))

	mux.Handle("/api/health", app.healthz())

	mux.Handle("/api/logs", app.viewLogsHandler())

	return mux
}
