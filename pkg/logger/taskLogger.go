package logger

import (
	"github.com/antoha2/task/config"
	amqp "github.com/rabbitmq/amqp091-go"
	//"github.com/streadway/amqp"
)

type Logger interface {
	Init() error
	Write(msg interface{})
}

type LoggerImpl struct {
	config      *config.Config
	amqpChannel *amqp.Channel
}

func NewLogger(cfg *config.Config) *LoggerImpl {

	return &LoggerImpl{config: cfg}
}

type LoggerMsg struct {
	Log string `json:"log"`
}
