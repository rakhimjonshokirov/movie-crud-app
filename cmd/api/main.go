package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"rakhimjonshokirov/movie-crud-app/internal/config"
	"rakhimjonshokirov/movie-crud-app/internal/gateways/rest"
	"rakhimjonshokirov/movie-crud-app/internal/infrastructure"
	"rakhimjonshokirov/movie-crud-app/internal/repositories"
	"rakhimjonshokirov/movie-crud-app/internal/usecases"
	"rakhimjonshokirov/movie-crud-app/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	_ "rakhimjonshokirov/movie-crud-app/api/docs"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	return router
}

func RegisterHooks(lc fx.Lifecycle, h *rest.Handler, cfg config.Config) {
	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: h,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Printf("Error starting server: %v\n", err)
					os.Exit(1)
				}
			}()
			fmt.Printf("Server started on %s\n", cfg.HTTPPort)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down server...")
			ctxShutDown, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return server.Shutdown(ctxShutDown)
		},
	})
}

func provideContext() context.Context {
	return context.Background()
}

// @title			movie-crud-app service API
// @version			1.0
// @description		Contains movie crud logics
// @contact.name	API Support
// @contact.url		https://www.linkedin.com/in/rakhimjonshokirov/
// @contact.email 	shokirovrakhimjon@gmail.com
// @schemes 		http
func main() {
	app := fx.New(
		// Provide constructors
		fx.Provide(
			provideContext,
			NewServer,
		),
		logger.Module,
		config.Module,
		infrastructure.Module,
		repositories.Module,
		usecases.Module,
		rest.Module,

		// Invoke starts up the server and any other initialization
		fx.Invoke(RegisterHooks),
	)

	app.Run()
}
