package server

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	machineryCfg "github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
)

func NewServer() *machinery.Server {
	cfg := &machineryCfg.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Redis: &machineryCfg.RedisConfig{
			MaxIdle:               3,
			IdleTimeout:           240,
			ReadTimeout:           15,
			WriteTimeout:          15,
			ConnectTimeout:        15,
			NormalTasksPollPeriod: 1000,
		},
	}

	broker := redisbroker.New(cfg, "localhost:6379", "", "", 0)
	backend := redisbackend.New(cfg, "localhost:6379", "", "", 0)
	lock := eagerlock.New()

	return machinery.NewServer(cfg, broker, backend, lock)
}
