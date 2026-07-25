package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func New() {
	// Create the storage log directory if it doesn't exist.
	if err := os.MkdirAll("storage/logs", 0755); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(
		"storage/logs/app.log",
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
