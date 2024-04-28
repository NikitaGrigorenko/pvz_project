package metrics

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ListAttemptMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "list_total_number_of_attempts",
		Help: "Total number of list command attempts.",
	})

	DeleteAttemptMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "delete_total_number_of_attempts",
		Help: "Total number of delete command attempts.",
	})

	AddAttemptMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "add_total_number_of_attempts",
		Help: "Total number of add command attempts.",
	})

	UpdateAttemptMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "update_total_number_of_attempts",
		Help: "Total number of list command attempts.",
	})

	GetAttemptMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "get_total_number_of_attempts",
		Help: "Total number of get command attempts.",
	})

	TotalNumberOfPVZMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_number_of_pvz",
		Help: "Total number of PVZ.",
	})
)

func RegisterMetrics() *prometheus.Registry {
	promMetrics := prometheus.NewRegistry()

	grpcMetrics := grpc_prometheus.NewServerMetrics()

	promMetrics.MustRegister(grpcMetrics, ListAttemptMetric, DeleteAttemptMetric, AddAttemptMetric, UpdateAttemptMetric, GetAttemptMetric, TotalNumberOfPVZMetric)

	return promMetrics
}
