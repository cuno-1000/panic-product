package serv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cuno-1000/panic-product/event/domain/model"
)

func init() {
	addr, err := primitive.NewNamesrvAddr("127.0.0.1:9876")
	if err != nil {
		panic(err)
	}
	Producer, err = rocketmq.NewProducer(
		producer.WithGroupName("EVENT"),
		producer.WithNameServer(addr),
		producer.WithCreateTopicKey(Topic),
		producer.WithRetry(1))
	if err != nil {
		panic(err)
	}

	err = Producer.Start()
	if err != nil {
		panic(err)
	}
}

var Producer rocketmq.Producer

const (
	Topic = "test"
)

func (e *EventDataService) PushToMq(record *model.EventApplyRecord) bool {
	message, err := json.Marshal(record)
	if err != nil {
		return false
	}

	err = Producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			fmt.Printf("receive message error:%v\n", err)
		} else {
			fmt.Printf("send message success. result=%s\n", result.String())
		}
	}, primitive.NewMessage(Topic, message))
	if err != nil {
		fmt.Printf("send async message error:%s\n", err)
		return false
	}
	return true
	//if res, err := Producer.SendSync(context.Background(), primitive.NewMessage(
	//	Topic,
	//	message,
	//)); err != nil {
	//	fmt.Printf("send sync message error:%s\n", err)
	//} else {
	//	fmt.Printf("send sync message success. result=%s\n", res.String())
	//}
}
