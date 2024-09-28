package log

import (
	"log/slog"
	"os"
)

func Default() *slog.Logger {
	opts := slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)
	return logger
}
