// https://youtu.be/kgkQZnh7BbI
package main

import (
	"log/slog"
	"os"
)

func main() {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler)

	// logger.Debug("What's the meaning of life?", "answer", 42) // smelly code if you have a lot of keys
	logger.Debug("What's the meaning of life?", slog.Int("answer", 42)) // smelly code if you have a lot of keys

	/*
		logger.Debug("debug level")
		logger.Info("test level") // default log is info
		logger.Warn("warn level")
		logger.Error("error level")
	*/
}
