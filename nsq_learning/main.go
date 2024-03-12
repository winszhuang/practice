package main

import (
	"go.uber.org/zap"
	"nsq_learning/mq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	m, err := mq.NewMessageQueue(mq.MessageQueueConfig{
		NsqAddr:         "127.0.0.1:4150",
		NsqLookupdAddr:  "127.0.0.1:4161",
		SupportedTopics: []string{"hello"},
	})

	if err != nil {
		zap.L().Fatal("Message queue error: " + err.Error())
	}

	m.Sub("hello", func(resp []byte) {
		zap.L().Info("S1 Got: " + string(resp))
	})
	m.Sub("hello", func(resp []byte) {
		zap.L().Info("S2 Got: " + string(resp))
	})
	m.Run()
	err = m.Pub("hello", "world")
	if err != nil {
		zap.L().Fatal("Message queue error: " + err.Error())
	}
	err = m.Pub("hello", "tom")
	if err != nil {
		zap.L().Fatal("Message queue error: " + err.Error())
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	os.Exit(0)
}

//
//const (
//	topic   = "topic"
//	channel = "channel"
//)
//
//var (
//	producer *nsq.Producer
//	consumer *nsq.Consumer
//)
//
//type myMessageHandler struct{}
//
//func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
//	s := string(m.Body)
//	fmt.Println("-----------------")
//	fmt.Println(s)
//	return nil
//}
//
//func main() {
//	initConsumer()
//	initProducer()
//	//go func() {
//	//}()
//	//
//	//go func() {
//	//}()
//
//	time.Sleep(2 * time.Second)
//	err := producer.Publish(topic, []byte("hello world"))
//	if err != nil {
//		panic(err)
//	}
//
//	time.Sleep(7 * time.Second)
//	producer.Stop()
//	consumer.Stop()
//}
//
//func initProducer() {
//	c := nsq.NewConfig()
//	var err error
//	producer, err = nsq.NewProducer("127.0.0.1:4150", c)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func initConsumer() {
//	c := nsq.NewConfig()
//	var err error
//	consumer, err = nsq.NewConsumer(topic, channel, c)
//	if err != nil {
//		panic(err)
//	}
//
//	consumer.AddHandler(&myMessageHandler{})
//
//	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
//	if err != nil {
//		panic(err)
//	}
//}
