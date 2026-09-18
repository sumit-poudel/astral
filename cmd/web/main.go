package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sumit-poudel/astral/internal/db"
)

type application struct {
	logger  *slog.Logger
	queries *db.Queries
}

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Error("Unable to connect to database: ", slog.String("pg", err.Error()))
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// server
	app := &application{
		logger:  logger,
		queries: db.New(conn),
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
