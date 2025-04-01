package main

import (
	"net/http"
	"os"
)

func (app *application) showOpenapiHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("docs/openapi.yaml")
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	w.WriteHeader(200)
	w.Write(f)
}
