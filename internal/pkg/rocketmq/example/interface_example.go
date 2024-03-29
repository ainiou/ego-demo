// Code generated by struct2interface; DO NOT EDIT.

package example

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// ConsumerExampleInterface ...
type ConsumerExampleInterface interface {
	// RegisterConsumerExample ...
	RegisterConsumerExample(f func(context.Context, *primitive.MessageExt) error)
}
type ProducerExampleInterface interface {
	// Send ...
	Send(ctx context.Context, data string) error
}
