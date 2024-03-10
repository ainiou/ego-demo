package example

import (
	"context"
	"ego-demo/pkg/erocketmq"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type ProducerExample struct {
	topic    string
	producer rocketmq.Producer
}

func NewProducerExampleInterface() ProducerExampleInterface {
	p, topic := erocketmq.LoadProducer("rocketmq.example.producer").BuildProducer()
	return &ProducerExample{
		topic:    topic,
		producer: p,
	}
}

// Send ...
func (e *ProducerExample) Send(ctx context.Context, data string) error {
	msg := &primitive.Message{
		Topic: e.topic,
		Body:  []byte(data),
	}

	// 同步发送消息
	resp, err := e.producer.SendSync(ctx, msg)
	if err != nil {
		return err
	}

	if resp.Status != primitive.SendOK {
		err = fmt.Errorf("send message failed: %d", resp.Status)
		return err
	}

	return nil
}
