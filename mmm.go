package main

import (
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int `env:"MMMDEBUG" env-default:"1"`
}

func setEnv() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})

	var cfg ConfigDatabase
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Info().Msg("[MMM] Unable to read environment variables, continuing with default.")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
	}

	// Set log level
	switch cfg.Debug {
	case 0: // Silent
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case 1: // Info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case 2: // Full errors
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		break
	default: // Default to Info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	}
	log.Info().Msg("[MMM] Log Level: " + strconv.Itoa(cfg.Debug) + ".")
}

func main() {
	setEnv()
}
