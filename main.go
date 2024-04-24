package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wit-switch/assessment-tax/config"
	"github.com/wit-switch/assessment-tax/infrastructure"
	"github.com/wit-switch/assessment-tax/pkg/errorx"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: errorx.ReplaceAttr,
	}))
	slog.SetDefault(logger)

	cfg := config.GetConfig()

	dbClient, err := infrastructure.NewPostgresClient(context.Background(), cfg.PostgreSQL)
	if err != nil {
		slog.Error("[!] failed to connect postgres", slog.Any("err", err))
		os.Exit(1)
	}

	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		ctx := c.Request().Context()
		if dbErr := dbClient.Ping(ctx); dbErr != nil {
			return dbErr
		}

		return c.String(http.StatusOK, "OK")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	server := &http.Server{
		Addr:              cfg.Server.HTTPAddress(),
		Handler:           e,
		ReadHeaderTimeout: 30 * time.Second,
	}

	quit := make(chan os.Signal, 1)

	go func() {
		if svErr := e.StartServer(server); !errorx.Is(svErr, http.ErrServerClosed) {
			slog.Error("[!] failed to serve server", slog.Any("err", svErr))
			os.Exit(1)
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	{
		<-quit
		slog.Info("gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if svErr := e.Shutdown(ctx); svErr != nil {
			slog.Error("[!] failed to shutdown server", slog.Any("err", svErr))
		}

		slog.Info("close postgres connection")
		dbClient.Close()

		slog.Info("shutting down the server")
	}
}
