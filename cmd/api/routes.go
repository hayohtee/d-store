package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /v1/key/{key}", app.putHandler)
	mux.HandleFunc("GET /v1/key/{key}", app.getHandler)
	mux.HandleFunc("DELETE /v1/key/{key}", app.deleteHandler)

	return mux
}
