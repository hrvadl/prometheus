package name

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var nameProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "testapp_name_requests_total",
	Help: "Total requests made to the /hello/{name} endpoint",
})

var requestProcessing = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "testapp_name_requests_pending",
	Help: "Total requests made to the /hello/{name} endpoint pending",
})

var timeTaken = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "testapp_name_requests_time",
	Help: "Total time spend waiting for the /hello/{name} endpoint to respond",
}, []string{"path"})

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

type Handler struct {
	log *slog.Logger
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	requestProcessing.Inc()
	defer requestProcessing.Dec()
	time.Sleep(time.Second * 5)
	nameProcessed.Inc()
	name := r.PathValue("name")
	w.WriteHeader(http.StatusOK)
	timeTaken.WithLabelValues("/hello/{name}").Observe(time.Since(now).Seconds())
	if _, err := w.Write([]byte(fmt.Sprintf("hello %s :0", name))); err != nil {
		h.log.Error("Failed write response", "err", err)
	}
}
