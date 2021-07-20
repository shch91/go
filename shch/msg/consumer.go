package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 新建一个arama配置实例
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092","localhost:9093","localhost:9094"}, "testGroup",config)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}

	defer consumer.Close()
	consumer.Consume(context.Background(),[]string{"mytopic"}, )
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

}
