package erocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
)

// consumerConfig consumer config
type consumerConfig struct {
	EndPoints []string
	Namespace string
	Enable    bool
	Topic     string
	Group     string
}

// DefaultConsumerConfig 返回默认配置
func DefaultConsumerConfig() *consumerConfig {
	return &consumerConfig{
		Enable: true,
	}
}

func LoadConsumer(key string) *Container {
	c := DefaultContainer()
	if err := econf.UnmarshalKey(key, &c.consumerConfig); err != nil {
		c.logger.Panic("parse config error", elog.FieldErr(err), elog.FieldKey(key))
		return c
	}
	c.logger = c.logger.With(elog.FieldComponentName(key))
	c.name = key
	return c
}

func (c *Container) BuildConsumer() (p rocketmq.PushConsumer, topic string) {
	// 检测topic配置
	if len(c.consumerConfig.Topic) == 0 {
		c.logger.Panic("rocketmq topic is empty", elog.FieldKey(c.name))
		return
	}
	topic = c.consumerConfig.Topic
	var err error
	if p, err = rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver(c.consumerConfig.EndPoints)),
		consumer.WithConsumerModel(consumer.Clustering), // 默认集群模式
		consumer.WithNamespace(c.consumerConfig.Namespace),
		consumer.WithGroupName(c.consumerConfig.Group),
	); err != nil {
		c.logger.Panic("rocketmq new consumer error", elog.FieldErr(err), elog.FieldKey(c.name))
		return
	}
	return
}
