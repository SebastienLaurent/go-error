package main

import (
	"os"
	
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/SebastienLaurent/go-error/api"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "20060102 15:04:05.000"})

	err := api.Test()
	if err != nil {
		if errors.Is(err, api.ErrApiNotWorking) { 
			if tracer , ok := err.(stackTracer) ; ok {
				stack := tracer.StackTrace()
				log.Error().Msgf("%+v", stack)
			}
		}

		log.Fatal().Err(err).Msg("can't invoke api")
	}
}
