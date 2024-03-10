package rocketmq

import (
	"ego-demo/internal/pkg/rocketmq/example"
	"github.com/google/wire"
)

// ProviderSet is wire provider set for wire
var ProviderSet = wire.NewSet(
	example.NewProducerExampleInterface,
	example.NewConsumerExampleInterface,
)
