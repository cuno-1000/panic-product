package handler

//func init() {
//	addr, err := primitive.NewNamesrvAddr("127.0.0.1:9876")
//	if err != nil {
//		panic(err)
//	}
//	Producer, err = rocketmq.NewProducer(
//		producer.WithGroupName("EVENT"),
//		producer.WithNameServer(addr),
//		producer.WithCreateTopicKey(Topic),
//		producer.WithRetry(1))
//	if err != nil {
//		panic(err)
//	}
//
//	err = Producer.Start()
//	if err != nil {
//		panic(err)
//	}
//
//}
//
//var Producer rocketmq.Producer
//
//const (
//	Topic = "test"
//)
//
//func PushToMq(requestParam *ApplyRequest) bool {
//	message, err := json.Marshal(requestParam)
//	if err != nil {
//		return false
//	}
//
//	err = Producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
//		if err != nil {
//			fmt.Printf("receive message error:%v\n", err)
//		} else {
//			fmt.Printf("send message success. result=%s\n", result.String())
//		}
//	}, primitive.NewMessage(Topic, message))
//	if err != nil {
//		fmt.Printf("send async message error:%s\n", err)
//		return false
//	}
//	return true
//}
//
//func ConsumeRequest() {
//	c, _ := rocketmq.NewPushConsumer(
//		consumer.WithGroupName("EVENT"),
//		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
//	)
//	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
//		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
//		for _, v := range msgs {
//			temp := &ApplyRequest{}
//			err := json.Unmarshal(v.Body, temp)
//			if err != nil {
//				return 0, err
//			}
//			MqProcess(temp.UserId, temp.EventUuid, temp.Link, temp.RequestId)
//		}
//		//这个相当于消费者 消息ack，如果失败可以返回 consumer.ConsumeRetryLater
//		return consumer.ConsumeSuccess, nil
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	// Note: start after subscribe
//	err = c.Start()
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(-1)
//	}
//}
