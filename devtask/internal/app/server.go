package app

import (
	"context"
	"devtask/internal/app/event"
	"devtask/internal/app/handlers"
	"devtask/internal/app/middleware"
	"devtask/internal/config"
	"devtask/internal/infrastructure/kafka"
	"devtask/internal/service/pvz"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const unsecurePort = ":9000"
const securePort = ":9001"

type Server struct {
	Repo handlers.StoragePVZ
}

func RunHTTP(_ context.Context, service *pvz.Service, config config.Config, sender *event.KafkaSender, recvr *kafka.Client, prom *prometheus.Registry) {
	implementation := Server{Repo: service}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	router := CreateRouter(implementation.Repo, prom)
	handler := middleware.BasicAuth(middleware.Logger(router, sender), config.AuthInfo.Username, config.AuthInfo.Password)

	unsecureServer := &http.Server{
		Addr:    unsecurePort,
		Handler: handler,
	}

	secureServer := &http.Server{
		Addr:    securePort,
		Handler: handler,
	}

	go func() {
		if err := secureServer.ListenAndServeTLS("server.crt", "server.key"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, " secure")
		}
	}()

	go func() {
		if err := unsecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, " unsecure")
		}
	}()

	go func() {
		kafka.ConsumerGroupManager(*recvr, config.Topic)
	}()

	<-quit

	fmt.Println("Завершение работы сервера!")

	err := unsecureServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err, unsecureServer)
	}

	err = secureServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err, unsecureServer)
	}

	fmt.Println("Работа сервера завершена!")
}

func CreateRouter(implementation handlers.StoragePVZ, prom *prometheus.Registry) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/pvz", handlers.Create(implementation)).Methods("POST")
	router.Handle("/pvz", handlers.List(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/pvz/{%s:[0-9]+}", handlers.QueryParamKey), handlers.GetByID(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/pvz/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Update(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/pvz/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Delete(implementation)).Methods("DELETE")
	router.Handle("/metrics", promhttp.HandlerFor(prom, promhttp.HandlerOpts{EnableOpenMetrics: true})).Methods("GET")
	return router
}
