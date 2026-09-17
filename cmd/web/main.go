package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	app := &application{
		logger: logger,
	}
	srv := &http.Server{
		Handler:  app.router(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Addr:     ":8090",
	}
	logger.Info("Striting server at ", slog.String("port", ":8090"))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Error: ", slog.String("err", err.Error()))
	}
}
