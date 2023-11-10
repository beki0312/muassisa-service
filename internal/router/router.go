package router

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
	"muassisa-service/internal/config"
	"muassisa-service/internal/handlers"
	"muassisa-service/internal/pkg/service"
	"net/http"
)

var EntryPoint = fx.Options(
	fx.Invoke(
		NewRouter,
	),
)

type dependencies struct {
	fx.In
	Lifecycle   fx.Lifecycle
	Config      config.IConfig
	SVC         service.IService
	Transaction handlers.ITransactionHandler
	//Worker      workers.IWorker
}

func NewRouter(d dependencies) {
	server := mux.NewRouter()
	mainRoute := server.PathPrefix("/api").Subrouter()
	routeVer := mainRoute.PathPrefix("/v1").Subrouter()
	transactionRoute := routeVer.PathPrefix("/muassisa").Subrouter()
	log.Println(transactionRoute)
	transactionRoute.HandleFunc("/get-language", d.Transaction.GetLanguage()).Methods("GET", "OPTIONS")
	transactionRoute.HandleFunc("/get-course", d.Transaction.GetCourse()).Methods("GET", "OPTIONS")
	srv := http.Server{
		Addr:    d.SVC.ConfigInstance().GetString("api.server.port"),
		Handler: server,
	}
	d.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)
}
