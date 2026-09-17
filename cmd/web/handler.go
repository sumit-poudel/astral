package main

import (
	"net/http"

	"github.com/sumit-poudel/astral/views"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	views.InderRenderer().Render(r.Context(), w)
}

func (app *application) shop(w http.ResponseWriter, r *http.Request) {
	views.ShopRenderer().Render(r.Context(), w)
}
