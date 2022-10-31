package main

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/repo"
	"github.com/cuno-1000/panic-product/repayment_record/domain/serv"
	"github.com/cuno-1000/panic-product/repayment_record/handler"
	repayment "github.com/cuno-1000/panic-product/repayment_record/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	ServiceName = "go.micro.srv.repayment.record"
)

func main() {
	listener, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Fatal(err)
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

	//err = db.AutoMigrate(&model.User{}, &model.Record{}, &model.BlackList{}).Error
	//if err != nil {
	//	return
	//}
	//rp1 := repo.NewRecordRepository(db)
	//err = rp1.InitTable()
	//rp2 := repo.NewBlacklistRepository(db)
	//err = rp2.InitTable()
	recordDataService := serv.NewRecordDataService(repo.NewRecordRepository(db))
	blacklistDataService := serv.NewBlacklistDataService(repo.NewBlacklistRepository(db))

	repayment.RegisterRepaymentRecordServer(grpcServer, &handler.Record{
		RecordDataService:    recordDataService,
		BlacklistDataService: blacklistDataService,
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
