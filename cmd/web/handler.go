package main

import (
	"net/http"

	"github.com/starfederation/datastar-go/datastar"
	"github.com/sumit-poudel/astral/views"
	"github.com/sumit-poudel/astral/views/templates"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	views.InderRenderer().Render(r.Context(), w)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.Login())
}

func (app *application) createAcc(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.CreateAccount())
}
