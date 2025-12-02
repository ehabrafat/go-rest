package main

import (
	"log/slog"
	"os"
)

func main() {

	application := &application{
		config: config{
			addr: ":9090",
		},
	}

	r := mount()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("App starting...")
	run(application, r)
}
