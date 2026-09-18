package main

import (
	"log/slog"
	"net/http"

	"github.com/sumit-poudel/astral/views"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	views.InderRenderer().Render(r.Context(), w)
}

func (app *application) shop(w http.ResponseWriter, r *http.Request) {
	users, err := app.queries.ListUsers(r.Context())
	if err != nil {
		app.logger.Debug("kun error", slog.String("err", err.Error()))
		w.Write([]byte("noting to show"))
		return
	}

	views.ShopRenderer(users).Render(r.Context(), w)
}
