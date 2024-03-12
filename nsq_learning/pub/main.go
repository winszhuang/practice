package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
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

	r.GET("/call", func(c *gin.Context) {
		message := "Hello, NSQ!" // 這是將要發送的消息

		// 發布消息到名為 `example_topic` 的主題
		err := producer.Publish("topic", []byte(message))
		if err != nil {
			log.Println("Could not publish message to NSQ", err)
			c.JSON(500, gin.H{"status": "error", "message": "Failed to publish message"})
			return
		}

		// 回傳成功的響應
		c.JSON(200, gin.H{"status": "success", "message": "Message published"})
	})

	r.Run(":8080")

	defer producer.Stop()
}
