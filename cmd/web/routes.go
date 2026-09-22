package main

import (
	"net/http"

	"github.com/sumit-poudel/astral/views"
)

func (app *application) router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(views.Files)))
	mux.HandleFunc("GET /", app.index)
	mux.HandleFunc("GET /login", app.login)
	mux.HandleFunc("GET /create", app.createAcc)
	return mux
}
