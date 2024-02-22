package main

import (
	"context"
	"time"

	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gofiber/fiber/v2"
	server "github.com/ngtrdai197/go-machinery/pkg/machinery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		initTasks()
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

func initTasks() {
	s := server.NewServer()
	helloTask := tasks.Signature{
		Name: "hello",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "world",
			},
		},
	}
	asyncResult, err := s.SendTaskWithContext(context.Background(), &helloTask)
	if err != nil {
		log.Err(err).Msg("An error occurred while sending a task")
		return
	}
	_, err = asyncResult.Get(5 * time.Second)
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
