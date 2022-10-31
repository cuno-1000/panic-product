package main

import (
	"github.com/cuno-1000/panic-product/user/domain/repo"
	"github.com/cuno-1000/panic-product/user/domain/serv"
	"github.com/cuno-1000/panic-product/user/handler"
	user "github.com/cuno-1000/panic-product/user/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	{
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
	// Database
	db, err := database()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	//rp := repo.NewUserRepository(db)
	//err = rp.InitTable()

	userDataService := serv.NewUserDataService(repo.NewUserRepository(db))
	bznDataService := serv.NewBznDataService(repo.NewUserRepository(db))

	user.RegisterUserServer(grpcServer, &handler.User{
		UserDataService: userDataService,
		BznDataService:  bznDataService,
	})

	if err = grpcServer.Serve(listener); err != nil {
		return
	}
}

type DbConnect struct {
	User     string
	Pwd      string
	Address  string
	Database string
}

func database() (*gorm.DB, error) {
	var mysqlInfo = &DbConnect{
		User:     "root",
		Pwd:      "admin",
		Address:  "127.0.0.1:3306",
		Database: "micro-2022",
	}
	return gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Address+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
}
