package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wit-switch/assessment-tax/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cfg := config.GetConfig()

	e := echo.New()

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
		if err := e.StartServer(server); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("[!] failed to serve server", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	{
		<-quit
		slog.Info("gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			slog.Error("[!] failed to shutdown server", slog.Any("err", err))
		}

		slog.Info("shutting down the server")
	}
}
