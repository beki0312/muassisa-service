package handlers

import (
	"go.uber.org/fx"
	"muassisa-service/internal/pkg/logger"
	"muassisa-service/internal/pkg/service"
	"net/http"
)

var NewHandler = fx.Provide(newHandler)

type IHandler interface {
	GetLanguage() http.HandlerFunc
	GetCourse() http.HandlerFunc
}
type dependencies struct {
	fx.In
	SVC    service.IService
	Logger logger.ILogger
}
type Handler struct {
	svc    service.IService
	Logger logger.ILogger
}

func newHandler(d dependencies) IHandler {
	return Handler{
		svc:    d.SVC,
		Logger: d.Logger,
	}
}
