package main

import (
	"net/http"
	"runtime/debug"

	"github.com/sumit-poudel/astral/cmd/web/model"
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

func (app *application) newTemplateData(r *http.Request) model.TemplateData {
	if app.isAuthenticated(r) {
		return model.TemplateData{
			UserName:        r.Context().Value(userNameConkextKey).(string),
			IsAuthenticated: true,
		}
	}
	return model.TemplateData{
		IsAuthenticated: false,
	}
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}
