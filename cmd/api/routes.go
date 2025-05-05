package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("PUT /v1/key/{key}", app.putHandler)
	
	return mux
}