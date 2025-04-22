package logger

import (
	"log/slog"
	"os"
	"strings"
)

func Init(env string) *slog.Logger {
	env = strings.ToLower(env)

	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
