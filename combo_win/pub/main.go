package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/action", func(c *gin.Context) {
		go action(producer)

		// 回傳成功的響應
		c.JSON(200, gin.H{"status": "success", "message": "Message published"})
	})

	r.Run(":8080")

	defer producer.Stop()
}

func action(producer *nsq.Producer) {
	messages := []string{
		"test-lose",
		"test-lose",
		"test-win",
		"test-win",
		"test-win",
		"test-lose",
		"test-win",
		"test-lose",
		"test-win",
		"test-lose",
		"test-lose",
		"test-win",
		"test-lose",
		"test-win",
		"test-lose",
	}

	go func() {
		for _, message := range messages {
			err := producer.Publish("topic", []byte(message))
			if err != nil {
				log.Println("Could not publish message to NSQ", err)
				return
			}
		}

	}()

	messages2 := []string{
		"qa-lose",
		"qa-lose",
		"qa-win",
		"qa-win",
		"qa-lose",
		"qa-win",
		"qa-win",
		"qa-lose",
		"qa-lose",
		"qa-lose",
		"qa-win",
		"qa-lose",
		"qa-win",
		"qa-lose",
	}

	go func() {
		for _, mess := range messages2 {
			err := producer.Publish("topic", []byte(mess))
			if err != nil {
				log.Println("Could not publish message to NSQ", err)
				return
			}
		}
	}()

	//
	//for i := 0; i < 10; i++ {
	//	go func(num int) {
	//		var message string
	//		if num%2 == 0 {
	//			message = fmt.Sprintf("win-%d", num)
	//		} else {
	//			message = fmt.Sprintf("lose-%d", num)
	//		}
	//		//message := genResult()
	//
	//		// 發布消息到名為 `example_topic` 的主題
	//		err := producer.Publish("topic", []byte(message))
	//		if err != nil {
	//			log.Println("Could not publish message to NSQ", err)
	//			c.JSON(500, gin.H{"status": "error", "message": "Failed to publish message"})
	//			return
	//		}
	//	}(i)
	//}
}

func genResult() string {
	if rand.Intn(2) == 0 {
		return "win"
	}
	return "lose"
}
