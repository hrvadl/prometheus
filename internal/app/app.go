package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hrvadl/prometheus/internal/transport/handlers/name"
)

func New(port int, log *slog.Logger) *App {
	return &App{
		port: port,
		log:  log,
	}
}

type App struct {
	port int
	log  *slog.Logger
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	r := chi.NewRouter()
	nh := name.NewHandler(a.log.With("source", "nameHandler"))
	r.Get("/hello/{name}", nh.Handle)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	return http.ListenAndServe(fmt.Sprintf(":%d", a.port), r)
}
