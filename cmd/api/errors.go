package main

import (
	"log/slog"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err.Error(), slog.String("request_method", r.Method), slog.String("request_url", r.URL.String()))
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	if err := writeJSON(w, status, env, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and cannot process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource cannot be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}