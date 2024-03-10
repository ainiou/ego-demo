package example

import (
	"context"
	"ego-demo/pkg/erocketmq"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gotomicro/ego/core/elog"
)

type ConsumerExample struct {
	topic    string
	consumer rocketmq.PushConsumer
}

func NewConsumerExampleInterface() ConsumerExampleInterface {
	p, topic := erocketmq.LoadConsumer("rocketmq.example.consumer").BuildConsumer()
	return &ConsumerExample{
		topic:    topic,
		consumer: p,
	}
}

// RegisterConsumerExample ...
func (c *ConsumerExample) RegisterConsumerExample(f func(context.Context, *primitive.MessageExt) error) {
	if err := c.consumer.Subscribe(c.topic, consumer.MessageSelector{}, func(ctx context.Context, msgList ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgList {
			err := f(ctx, msg)
			if err != nil {
				return consumer.ConsumeRetryLater, err
			}
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		elog.Panic("rocketmq consumer subscribe error", elog.FieldErr(err), elog.FieldKey(c.topic))
		return
	}

	if err := c.consumer.Start(); err != nil {
		elog.Panic("rocketmq consumer start error", elog.FieldErr(err), elog.FieldKey(c.topic))
		return
	}
}
