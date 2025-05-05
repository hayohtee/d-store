package main

import (
	"net/http"

	"github.com/hayohtee/d-store/internal/storage"
)

func (app *application) putHandler(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	var payload struct {
		Value string `json:"value"`
	}

	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := storage.Put(key, payload.Value); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, envelope{"message": "item created successfully"}, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
