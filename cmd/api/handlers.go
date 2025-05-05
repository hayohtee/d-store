package main

import (
	"net/http"

	"github.com/hayohtee/d-store/internal/storage"
	"github.com/hayohtee/d-store/internal/validator"
)

func (app *application) putHandler(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	var payload struct {
		Value string `json:"value"`
	}

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(payload.Value != "", "value", "must be provided")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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
