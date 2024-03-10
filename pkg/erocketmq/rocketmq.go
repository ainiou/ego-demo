package erocketmq

import (
	"github.com/gotomicro/ego/core/elog"
)

// Container ...
type Container struct {
	producerConfig *producerConfig
	consumerConfig *consumerConfig
	name           string
	logger         *elog.Component
}

// DefaultContainer ...
func DefaultContainer() *Container {
	return &Container{
		producerConfig: DefaultProducerConfig(),
		consumerConfig: DefaultConsumerConfig(),
		logger:         elog.EgoLogger.With(elog.FieldComponent("component.rmq")),
	}
}
