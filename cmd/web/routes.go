package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/sumit-poudel/astral/views"
)

func (app *application) router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(views.Files)))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.authenticate)
	mux.Handle("GET /", dynamic.ThenFunc(app.getIndex))
	mux.Handle("GET /login", dynamic.ThenFunc(app.getLogin))
	mux.Handle("POST /login", dynamic.ThenFunc(app.postLogin))
	mux.Handle("GET /signup", dynamic.ThenFunc(app.getSignup))
	mux.Handle("POST /signup", dynamic.ThenFunc(app.postSignup))

	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("POST /logout", protected.ThenFunc(app.logOut))
	mux.Handle("GET /test", protected.ThenFunc(app.test))

	return mux
}
