package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) clientError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)
	app.logger.Error(err.Error(), "URL", uri, "method", method, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) decodeForm(r *http.Request, dest any) (err error) {
	if err := r.ParseForm(); err != nil {
		return err
	}
	if err := app.formDecoder.Decode(dest, r.PostForm); err != nil {
		return err
	}
	return nil
}
