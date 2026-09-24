package main

import (
	"context"
	"net/http"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")
		name := app.sessionManager.GetString(r.Context(), "authenticatedUserName")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}

		exists, err := app.queries.Exists(r.Context(), id)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		if exists {
			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
			ctx = context.WithValue(ctx, userIdConkextKey, id)
			ctx = context.WithValue(ctx, userNameConkextKey, name)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
