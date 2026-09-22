package main

import (
	"net/http"

	"github.com/sumit-poudel/astral/views"
)

func (app *application) router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(views.Files)))
	mux.HandleFunc("GET /", app.getIndex)
	mux.HandleFunc("GET /login", app.getLogin)
	mux.HandleFunc("POST /login", app.postLogin)
	mux.HandleFunc("GET /signup", app.getSignup)
	mux.HandleFunc("POST /signup", app.postSignup)
	return app.sessionManager.LoadAndSave(mux)
}
