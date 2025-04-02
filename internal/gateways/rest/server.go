package rest

import (
	"net/http"
	"rakhimjonshokirov/movie-crud-app/internal/config"

	"github.com/gin-gonic/gin"
	"gitlab.hamkorbank.uz/libs/logger/instrumentation/ginlog"
	"go.uber.org/zap"
)

var _whiteListHeader = []string{"*"}

type Handler struct {
	router   *gin.Engine
	movieUsc movieUscInterface
	usersUsc usersUscInterface
	l        *zap.Logger
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func newHandler(cfg config.Config, l *zap.Logger, r *gin.Engine,
	m movieUscInterface, u usersUscInterface) *Handler {
	baseRouter := r
	baseRouter.Use(ginlog.Log(), ginlog.RecoveryWithZap(false))

	h := Handler{
		l:        l,
		router:   baseRouter,
		movieUsc: m,
		usersUsc: u,
	}
	h.routes()

	return &h
}
