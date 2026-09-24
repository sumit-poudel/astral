package main

import (
	"net/http"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
	"github.com/sumit-poudel/astral/internal/db"
	"github.com/sumit-poudel/astral/internal/validator"
	"github.com/sumit-poudel/astral/views"
	"github.com/sumit-poudel/astral/views/templates"
)

func (app *application) logOut(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.sessionManager.Remove(r.Context(), "authenticatedUserName")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is protected"))
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	views.InderRenderer(data).Render(r.Context(), w)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.Login())
	sse.PatchSignals([]byte(`{"modalOpen": true}`))
}

func (app *application) getSignup(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(templates.Signup())
}

// signup ko lagi
// signup ko handler k
type userSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) postSignup(w http.ResponseWriter, r *http.Request) {
	signupForm := &userSignupForm{}

	if err := app.decodeForm(r, signupForm); err != nil {
		app.clientError(w, http.StatusUnprocessableEntity)
		return
	}
	// validation suru vayo yo bata
	signupForm.CheckField(signupForm.NotBlank(signupForm.Email), "Email", "email cant be empty!")
	signupForm.CheckField(signupForm.Matches(signupForm.Email, validator.EmailRx), "Email", "Email address not valid!")
	signupForm.CheckField(signupForm.NotBlank(signupForm.Name), "Name", "name cant be empty!")
	signupForm.CheckField(signupForm.NotBlank(signupForm.Password), "Password", "password cant be empty!")
	if !signupForm.IsValid() {
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(templates.Error(signupForm.FieldErrors))
		return
	}
	if err := signupForm.CheckFieldWithErr(func() (bool, error) {
		found, err := app.queries.EmailTaken(r.Context(), signupForm.Email)
		return !found, err
	}, "Email", "Email already taken!"); err != nil {
		app.serverError(w, r, err)
		return
	}
	if !signupForm.IsValid() {
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(templates.Error(signupForm.FieldErrors))
		return
	}
	// after validation
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

// login ko lagi
// login ko handler k
type userLoginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
	loginForm := &userLoginForm{}
	data := &db.LoginUserRow{}
	if err := app.decodeForm(r, loginForm); err != nil {
		app.clientError(w, http.StatusUnprocessableEntity)
		return
	}
	loginForm.CheckField(loginForm.NotBlank(loginForm.Email), "Email", "email cant be empty")
	loginForm.CheckField(loginForm.Matches(loginForm.Email, validator.EmailRx), "Email", "Email address not valid!")
	loginForm.CheckField(loginForm.NotBlank(loginForm.Password), "Password", "password cant be empty")
	if !loginForm.IsValid() {
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(templates.Error(loginForm.FieldErrors))
		return
	}
	if err := loginForm.CheckFieldWithErr(func() (bool, error) {
		user, err := app.queries.LoginUser(r.Context(), loginForm.Email)
		if err != nil {
			return false, err
		}
		if strings.Compare(user.Password, loginForm.Password) == 0 {
			data = &user
			return true, nil
		}
		return false, nil
	}, "Password", "password incorrect!"); err != nil {
		app.serverError(w, r, err)
		return
	}
	if !loginForm.IsValid() {
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(templates.Error(loginForm.FieldErrors))
		return
	}
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.sessionManager.Put(r.Context(), "authenticatedUserID", data.Uid)
	app.sessionManager.Put(r.Context(), "authenticatedUserName", data.Name)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
