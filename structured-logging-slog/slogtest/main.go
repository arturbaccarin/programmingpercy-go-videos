// https://youtu.be/kgkQZnh7BbI
package main

import (
	"log/slog"
	"os"
	"time"
)

// 20min
func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ // NewTextHandler
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr { // iterate over attributes
			// Match the key that we want
			if a.Key == slog.TimeKey {
				a.Key = "date"
				a.Value = slog.Int64Value(time.Now().Unix())
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.Int("answer", 42), // applications informations, versions
		slog.Group("vodes",
			slog.Int("Pikachu", 40), // system information inside system group
			slog.Int("Mew", 22),
		),
	})

	logger := slog.New(logHandler)

	logger.Debug("Best Pokemon Rating")

	slog.SetDefault(logger) // setup with the current logger

	slog.Info("New Info")

	// logger.Debug("What's the meaning of life?", "answer", 42) // smelly code if you have a lot of keys
	// logger.Debug("What's the meaning of life?", slog.Int("answer", 42)) // smelly code if you have a lot of keys

	/*
			logger.Debug("debug level")
			logger.Info("test level") // default log is info
			logger.Warn("warn level")
			logger.Error("error level")

			logger.Debug("Best Pokemon Rating",
			slog.Group("vodes",
				slog.Int("Pikachu", 40), // system information inside system group
				slog.Int("Mew", 22),
			),
		)
	*/
}
