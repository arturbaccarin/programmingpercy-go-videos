// https://youtu.be/kgkQZnh7BbI
package main

import (
	"log/slog"
	"os"
)

// 20min
func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ // NewTextHandler
		Level:     slog.LevelDebug,
		AddSource: true,
	}).WithAttrs([]slog.Attr{
		slog.Int("answer", 42), // applications informations, versions
	})

	logger := slog.New(logHandler)

	logger.Debug("Best Pokemon Rating",
		slog.Group("vodes",
			slog.Int("Pikachu", 40), // system information inside system group
			slog.Int("Mew", 22),
		),
	)

	// logger.Debug("What's the meaning of life?", "answer", 42) // smelly code if you have a lot of keys
	// logger.Debug("What's the meaning of life?", slog.Int("answer", 42)) // smelly code if you have a lot of keys

	/*
		logger.Debug("debug level")
		logger.Info("test level") // default log is info
		logger.Warn("warn level")
		logger.Error("error level")
	*/
}
