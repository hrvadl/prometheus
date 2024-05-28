package main

import (
	"log/slog"
	"os"

	"github.com/hrvadl/prometheus/internal/app"
)

const port = 5005

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := app.New(port, log)
	app.MustRun()
}
