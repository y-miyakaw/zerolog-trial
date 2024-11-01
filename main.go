package main

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logLevel zerolog.Level

func getLogLevel() zerolog.Level {
	if os.Getenv("DEBUG_LOG") == "true" {
		return zerolog.DebugLevel
	}
	return zerolog.InfoLevel
}

func HandleRequest(ctx context.Context, event interface{}) {
	zerolog.SetGlobalLevel(getLogLevel())

	log.Debug().Interface("event", event).Msg("Received event")
}

func main() {
	// Start the server
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		zerolog.SetGlobalLevel(getLogLevel())
		log.Debug().Msg("Hello, World!")
		log.Info().Msg("Hello, World!")
		log.Warn().Msg("Hello, World!")
		log.Error().Msg("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":8080")
}
