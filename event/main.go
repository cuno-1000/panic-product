package main

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/cuno-1000/panic-product/event/domain/repo"
	"github.com/cuno-1000/panic-product/event/domain/serv"
	"github.com/cuno-1000/panic-product/event/handler"
	event "github.com/cuno-1000/panic-product/event/proto"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func main() {
	listener, err := net.Listen("tcp", ":3456")
	if err != nil {
		log.Fatal(err)
		return
	}

	{
		serv.UserServiceConn, err = grpc.Dial("localhost:1234", grpc.WithInsecure())
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
		}(serv.UserServiceConn)

		serv.RepaymentServiceConn, err = grpc.Dial("localhost:2345", grpc.WithInsecure())
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
		}(serv.RepaymentServiceConn)
	}

	grpcServer := grpc.NewServer()

	redisDb := redisInit()

	eventDataService := serv.NewEventDataService(repo.NewEventRepository(redisDb))

	event.RegisterEventEngineServer(grpcServer, &handler.Event{
		EventDataService: eventDataService,
	})

	go eventDataService.WarmUpEventRoutines()

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
		return
	}
}

type RedisConnect struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func redisInit() *redis.Client {
	var redisInfo = &RedisConnect{
		Host:     "127.0.0.1",
		Port:     6379,
		DB:       0,
		Password: "",
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Host + ":" + strconv.Itoa(redisInfo.Port),
		Password: redisInfo.Password, // no password set
		DB:       redisInfo.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	return client
}
