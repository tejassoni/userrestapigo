package utils

import (
	"io"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(
		"logs/app.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	handler := slog.NewTextHandler(
		io.MultiWriter(os.Stdout, file),
		&slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		},
	)

	Logger = slog.New(handler)

	// Make it the default logger
	slog.SetDefault(Logger)
}
