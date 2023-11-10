package handlers

import (
	"go.uber.org/fx"
	"muassisa-service/internal/pkg/logger"
	"muassisa-service/internal/pkg/service"
	"net/http"
)

var NewTransactionHandler = fx.Provide(newTransactionHandler)

type ITransactionHandler interface {
	GetLanguage() http.HandlerFunc
	GetCourse() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC    service.IService
	Logger logger.ILogger
}

type TransactionHandler struct {
	svc    service.IService
	Logger logger.ILogger
}

func newTransactionHandler(d dependencies) ITransactionHandler {
	return TransactionHandler{
		svc:    d.SVC,
		Logger: d.Logger,
	}
}
