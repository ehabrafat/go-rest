package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // BLANK IMPORT - registers driver
)

func main() {

	godotenv.Load()

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app := &application{
		config: config{
			addr: ":9090",
			dbConfig: dbConfig{
				conStr: "postgres://postgres:secret@localhost:5432/pgtest?sslmode=disable",
			},
		},
	}

	con, err := pgx.Connect(ctx, app.config.dbConfig.conStr)
	
	if err != nil {
		panic(err.Error())
	}

	defer con.Close(ctx)
	
	logger.Info("Connected to database")

	app.db = con

	r := app.mount()

	slog.Info("App starting...")

	run(app, r)
}
