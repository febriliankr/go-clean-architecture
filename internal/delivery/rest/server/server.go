package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/febriliankr/go-clean-architecture/internal/app/config"
	"github.com/febriliankr/go-clean-architecture/internal/delivery/rest/middleware"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

// New instantiates new Echo server
func New() *echo.Echo {
	e := echo.New()
	e.Use(md.CORSWithConfig(md.CORSConfig{
		AllowOrigins:     []string{"*"},
		MaxAge:           86400,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}), middleware.Headers())
	e.GET("/health", healthCheck)
	e.Pre(md.RemoveTrailingSlash())
	e.Validator = &CustomValidator{V: validator.New()}
	customErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = customErr.handler
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}
	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

// Start starts echo server
func Start(e *echo.Echo, cfg *config.Server) {
	s := &http.Server{
		Addr:         ":" + cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}
	e.Debug = cfg.Debug

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Fatal("shutting down the server. error: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 4)
	signal.Notify(quit, os.Interrupt)
	<-quit
	e.Logger.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
