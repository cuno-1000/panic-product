package serv

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/repo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"reflect"
	"testing"
	"time"
)

//func TestNewRecordDataService(t *testing.T) {
//	type args struct {
//		recordRepository repo.IRecordRepository
//	}
//	tests := []struct {
//		name string
//		args args
//		want IRecordDataService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewRecordDataService(tt.args.recordRepository); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewRecordDataService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

type DbConnect struct {
	User     string
	Pwd      string
	Address  string
	Database string
}

func TestRecordDataService_AddRecord(t *testing.T) {
	var mysqlInfo = &DbConnect{
		User:     "root",
		Pwd:      "admin",
		Address:  "127.0.0.1:3306",
		Database: "micro-2022",
	}
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Address+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	type fields struct {
		RecordRepository repo.IRecordRepository
	}
	type args struct {
		userId uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"good",
			fields{
				RecordRepository: repo.NewRecordRepository(db),
			},
			args{
				2,
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecordDataService{
				RecordRepository: tt.fields.RecordRepository,
			}
			if err := r.AddRecord(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("AddRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecordDataService_GetUserIDNotAcceptedToPanicEvent(t *testing.T) {
	var mysqlInfo = &DbConnect{
		User:     "root",
		Pwd:      "admin",
		Address:  "127.0.0.1:3306",
		Database: "micro-2022",
	}
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Address+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	type fields struct {
		RecordRepository repo.IRecordRepository
	}
	type args struct {
		timeUpperLimit time.Time
		exp            string
		maxTimes       uint
	}
	timeUpperLimit, err := time.Parse("2006-01-02 15:04:05", "2020-04-10 17:00:00")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []uint
		wantErr bool
	}{
		{
			"good",
			fields{
				RecordRepository: repo.NewRecordRepository(db),
			},
			args{
				timeUpperLimit,
				"date_difference < 3 AND owed_amount < 3000",
				2,
			},
			[]uint{2},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecordDataService{
				RecordRepository: tt.fields.RecordRepository,
			}
			got, err := r.GetUserIDNotAcceptedToPanicEvent(tt.args.timeUpperLimit, tt.args.exp, tt.args.maxTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserIDNotAcceptedToPanicEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserIDNotAcceptedToPanicEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordDataService_IsAcceptedToPanicEventByUserId(t *testing.T) {
	var mysqlInfo = &DbConnect{
		User:     "root",
		Pwd:      "admin",
		Address:  "127.0.0.1:3306",
		Database: "micro-2022",
	}
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Address+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	timeUpperLimit, err := time.Parse("2006-01-02 15:04:05", "2020-04-10 17:00:00")
	type fields struct {
		RecordRepository repo.IRecordRepository
	}
	type args struct {
		timeUpperLimit time.Time
		userId         uint64
		exp            string
		maxTimes       uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"good",
			fields{RecordRepository: repo.NewRecordRepository(db)},
			args{
				timeUpperLimit,
				7,
				"date_difference < 3 AND owed_amount < 3000",
				2,
			},
			true,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecordDataService{
				RecordRepository: tt.fields.RecordRepository,
			}
			got, err := r.IsAcceptedToPanicEventByUserId(tt.args.timeUpperLimit, tt.args.userId, tt.args.exp, tt.args.maxTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAcceptedToPanicEventByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAcceptedToPanicEventByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
