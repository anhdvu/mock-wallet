package main

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

func (app *application) viewLogsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			app.respondJSONWithError(w, http.StatusMethodNotAllowed, "only GET method is allowed")
			return
		}

		page, size := app.getURLParams(r)

		filter := logFilter{
			offset: size * (page - 1),
			limit:  size,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		logs, err := app.apiLogger.FindLogs(ctx, filter)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
			return
		}

		err = app.respondJSON(w, http.StatusOK, logs)
		if err != nil {
			app.logger.Println(err)
			app.serverErrorResponse(w, err)
		}
	}
}

func (app *application) getURLParams(r *http.Request) (int, int) {
	params := r.URL.Query()
	pageParam := params.Get("page")
	sizeParam := params.Get("size")

	var err error
	var page, size int

	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			app.logger.Println(err)
			page = 1
		}
	}

	if page < 1 {
		page = 1
	}

	if sizeParam != "" {
		size, err = strconv.Atoi(sizeParam)
		if err != nil || size < 1 {
			app.logger.Println(err)
			size = 10
		}
	}

	if size < 1 {
		size = 10
	}

	if size > 50 {
		size = 50
	}

	return page, size
}
