package mq

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"time"
)

type MessageQueueConfig struct {
	NsqAddr         string
	NsqLookupdAddr  string
	SupportedTopics []string
}

type MessageQueue struct {
	config    MessageQueueConfig
	producer  *nsq.Producer
	consumers map[string]*nsq.Consumer
}

func NewMessageQueue(config MessageQueueConfig) (mq *MessageQueue, err error) {
	zap.L().Debug("New message queue")
	producer, err := initProducer(config.NsqAddr)
	if err != nil {
		return nil, err
	}
	consumers := make(map[string]*nsq.Consumer)
	for _, topic := range config.SupportedTopics {
		nsq.Register(topic, "default")
		consumers[topic], err = initConsumer(topic, "default", config.NsqAddr)
		if err != nil {
			return
		}
	}
	return &MessageQueue{
		config:    config,
		producer:  producer,
		consumers: consumers,
	}, nil
}

func (mq *MessageQueue) Run() {
	for name, c := range mq.consumers {
		zap.L().Info("Run consumer for " + name)
		// c.ConnectToNSQLookupd(mq.config.NsqLookupdAddr)
		c.ConnectToNSQD(mq.config.NsqAddr)
	}
}

func initProducer(addr string) (producer *nsq.Producer, err error) {
	zap.L().Debug("initProducer to " + addr)
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(addr, config)
	return
}

func initConsumer(topic string, channel string, address string) (c *nsq.Consumer, err error) {
	zap.L().Debug("initConsumer to " + topic + "/" + channel)
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err = nsq.NewConsumer(topic, channel, config)
	return
}

func (mq *MessageQueue) Pub(name string, data interface{}) (err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}
	zap.L().Info("Pub " + name + " to mq. data = " + string(body))
	return mq.producer.Publish(name, body)
}

type Messagehandler func(v []byte)

func (mq *MessageQueue) Sub(name string, handler Messagehandler) (err error) {
	zap.L().Info("Subscribe " + name)
	v, ok := mq.consumers[name]
	if !ok {
		err = fmt.Errorf("No such topic: " + name)
		return
	}
	v.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		handler(message.Body)
		return nil
	}))
	return
}
