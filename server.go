package main

import (
	"github.com/rs/zerolog/log"

	"github.com/oscaralmgren/rest-banking-api-go-tutorial/app"
)

func main() {
	// zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	log.Info().Msg("Starting rest-banking-api-go-tutorial api")
	// log.Error().Msg("Error rest-banking-api-go-tutorial api")
	app.Start()
}
