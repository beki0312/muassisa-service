package router

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
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
	Lifecycle fx.Lifecycle
	//Config    config.IConfig
	SVC     service.IService
	Handler handlers.IHandler
}

func NewRouter(d dependencies) {
	server := mux.NewRouter()
	mainRoute := server.PathPrefix("/api").Subrouter()
	routeVer := mainRoute.PathPrefix("/v1").Subrouter()
	courseRoute := routeVer.PathPrefix("/muassisa").Subrouter()
	courseRoute.HandleFunc("/added_course", d.Handler.AddedCourse()).Methods("POST", "OPTIONS")
	courseRoute.HandleFunc("/get-language", d.Handler.GetLanguage()).Methods("GET", "OPTIONS")
	courseRoute.Path("/get-course").Queries("id", "{id}").HandlerFunc(d.Handler.GetCourse()).Methods("GET", "OPTIONS")
	courseRoute.HandleFunc("/get-courses", d.Handler.GetCourseNew()).Methods("GET", "OPTIONS")
	srv := http.Server{
		Addr:    "127.0.5.24:53488", //os.Getenv("APP_PORT"), //d.SVC.ConfigInstance().GetString("api.server.port"),
		Handler: server,
	}
	log.Println("srv  ", srv.Addr)
	d.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				//d.SVC.LoggerInstance().Info("Application started")
				log.Println("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				//d.SVC.LoggerInstance().Info("Application stopped")
				log.Println("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)
}
