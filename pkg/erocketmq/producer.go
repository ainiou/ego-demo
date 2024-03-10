package erocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
)

// producerConfig producer config
type producerConfig struct {
	EndPoints []string
	Namespace string
	Topic     string
	Group     string
	Retry     int // 重试次数
}

// DefaultProducerConfig 返回默认配置
func DefaultProducerConfig() *producerConfig {
	return &producerConfig{
		EndPoints: make([]string, 0),
		Retry:     5,
	}
}

func LoadProducer(key string) *Container {
	c := DefaultContainer()
	if err := econf.UnmarshalKey(key, &c.producerConfig); err != nil {
		c.logger.Panic("parse config error", elog.FieldErr(err), elog.FieldKey(key))
		return c
	}
	c.logger = c.logger.With(elog.FieldComponentName(key))
	c.name = key
	return c
}

func (c *Container) BuildProducer() (p rocketmq.Producer, topic string) {
	// 检测topic配置
	if len(c.producerConfig.Topic) == 0 {
		c.logger.Panic("rocketmq topic is empty", elog.FieldKey(c.name))
		return
	}
	topic = c.producerConfig.Topic

	var err error
	if p, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(c.producerConfig.EndPoints)),
		producer.WithRetry(c.producerConfig.Retry),
		producer.WithNamespace(c.producerConfig.Namespace),
		producer.WithGroupName(c.producerConfig.Group),
	); err != nil {
		c.logger.Panic("rocketmq new producer error", elog.FieldErr(err), elog.FieldKey(c.name))
		return
	}

	if err = p.Start(); err != nil {
		c.logger.Panic("rocketmq producer start error", elog.FieldErr(err), elog.FieldKey(c.name))
		return
	}

	return
}
