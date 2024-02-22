package main

import (
	"github.com/RichardKnop/machinery/v2/tasks"
	server "github.com/ngtrdai197/go-machinery/pkg/machinery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	consumerTag := "machinery_worker"
	server := server.NewServer()

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerTag, 0)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorHandler := func(err error) {
		log.Err(err).Msg("An error occurred while processing a task")
	}

	preTaskHandler := func(signature *tasks.Signature) {
		log.Info().Msgf("Pre state Task ID: %s", signature.UUID)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		log.Info().Msgf("PostState Task ID: %s", signature.UUID)
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)

	worker.Launch()
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
