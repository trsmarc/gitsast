package queue

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/go-redis/redis/v8"

	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

type Handler interface {
	StartConsumer() error
	AddTask(t *taskq.Message) error
}

type handler struct {
	redis        *redis.Client
	queueFactory taskq.Factory
	mainQueue    taskq.Queue
}

func NewHandler() Handler {
	redis := redis.NewClient(&redis.Options{
		Addr:       ":6379",
		MaxRetries: 3,
	})

	queueFactory := redisq.NewFactory()

	mainQueue := queueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "main-queue",
		Redis: redis,
	})

	return &handler{redis, queueFactory, mainQueue}
}

func (h *handler) AddTask(t *taskq.Message) error {
	return h.mainQueue.Add(t)
}

func (h *handler) StartConsumer() error {
	queueConsumerErr := make(chan error)

	go func() {
		queueConsumerErr <- h.queueFactory.StartConsumers(context.Background())
	}()

	shutdown := make(chan os.Signal, 1)

	select {
	case err := <-queueConsumerErr:
		return err
	case sig := <-shutdown:
		log.Info().Any("signal", sig).Msg("stopping queue consumer with signal")
		err := h.queueFactory.StopConsumers()
		if err != nil {
			return err
		}

		err = h.queueFactory.Close()
		if err != nil {
			return err
		}

		log.Info().Any("signal", sig).Msg("stopped queue consumer with signal")
	}
	return nil
}
