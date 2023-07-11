package main

import (
	"encoding/xml"
	"net/http"
)

func (app *application) respondXML(w http.ResponseWriter, status int, data any) error {
	payload, err := xml.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}
