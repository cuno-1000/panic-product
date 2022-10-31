package main

import (
	"github.com/cuno-1000/panic-product/api-gateway/handler"
	"github.com/cuno-1000/panic-product/api-gateway/routes"
	"google.golang.org/grpc"
	"log"
	"runtime"
)

func main() {
	// cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(4)

	var err error

	{
		handler.UserServiceConn, err = grpc.Dial("localhost:1234", grpc.WithInsecure())
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
		}(handler.UserServiceConn)

		handler.EventServiceConn, err = grpc.Dial("localhost:3456", grpc.WithInsecure())
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
		}(handler.EventServiceConn)
	}

	//go handler.ConsumeRequest()
	handler.ApplyPool = handler.NewWorkerPool(200)
	handler.ApplyPool.Run()

	router := routes.NewRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
		return
	}
}
