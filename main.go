package main

import (
	"log/slog"
	"os"

	"github.com/wit-switch/assessment-tax/cmd/rest"
	"github.com/wit-switch/assessment-tax/config"
	"github.com/wit-switch/assessment-tax/pkg/errorx"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: errorx.ReplaceAttr,
	}))
	slog.SetDefault(logger)

	cfg := config.GetConfig()

	rest.Execute(cfg)
}
