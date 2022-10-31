package main

import (
	"context"
	"fmt"
	pb "github.com/cuno-1000/panic-product/event/proto"
	"google.golang.org/grpc"
	"log"
)

var EventServiceConn *grpc.ClientConn

func Create() {
	var err error
	EventServiceConn, err = grpc.Dial("8.134.38.30:3456", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(EventServiceConn)

	client := pb.NewEventEngineClient(EventServiceConn)
	rsp, err := client.CreateEvent(context.Background(), &pb.CreateEventRequest{
		Info:               "BBBBBBB",
		AdminId:            500,
		ApplyRules:         "date_difference < 3 AND owed_amount < 3000",
		ProductQuantity:    10000,
		ProductItemPrice:   1000,
		ReviewUpperLimitAt: "2020-04-10 17:00:00",
		OverDueMaxTimes:    2,
		StartingAt:         "2022-04-10 22:00:00",
	})
	if err != nil {
		log.Fatal("Connect Fail:", err)
	}
	if rsp.IsSuccess {
		fmt.Println("GOOD!")
	}
}

func main() {
	Create()
}
