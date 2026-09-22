package main

import (
	"net/http"

	"github.com/starfederation/datastar-go/datastar"
	"github.com/sumit-poudel/astral/internal/db"
	"github.com/sumit-poudel/astral/views"
	"github.com/sumit-poudel/astral/views/templates"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")
	views.InderRenderer(id).Render(r.Context(), w)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.Login())
}

func (app *application) getSignup(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.Signup())
}

type userSignupForm struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (app *application) postSignup(w http.ResponseWriter, r *http.Request) {
	signupForm := &userSignupForm{}
	if err := app.decodeForm(r, signupForm); err != nil {
		app.clientError(w, http.StatusUnprocessableEntity)
		return
	}
	id, err := app.queries.SignupUser(r.Context(), db.SignupUserParams{Name: signupForm.Name, Email: signupForm.Email, Password: signupForm.Password})
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	app.sessionManager.Put(r.Context(), "authenticatedUserName", signupForm.Name)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
}
