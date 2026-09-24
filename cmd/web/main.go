package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sumit-poudel/astral/internal/db"
)

type application struct {
	logger         *slog.Logger
	queries        *db.Queries
	sessionManager *scs.SessionManager
	formDecoder    *form.Decoder
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if err := godotenv.Load(); err != nil {
		logger.Warn(".env file not found")
	}

	dsn := os.Getenv("GOOSE_DBSTRING")

	pool, err := openDB(dsn)

	if err != nil {
		logger.Error("Unable to connect to database: ", slog.String("pg", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Store = pgxstore.New(pool)

	app := &application{
		logger:         logger,
		queries:        db.New(pool),
		sessionManager: sessionManager,
		formDecoder:    form.NewDecoder(),
	}
	srv := &http.Server{
		Handler:  app.router(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Addr:     ":8090",
	}
	logger.Info("Striting server at ", slog.String("port", ":8090"))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Error: ", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func openDB(dns string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return pool, err
}
