package serv

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/repo"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"testing"
)

func TestBlacklistDataService_AppendBlacklist(t *testing.T) {
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
		BlacklistRepository repo.IBlacklistRepository
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
				BlacklistRepository: repo.NewBlacklistRepository(db),
			},
			args{
				8,
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlacklistDataService{
				BlacklistRepository: tt.fields.BlacklistRepository,
			}
			if err := b.AppendBlacklist(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("AppendBlacklist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlacklistDataService_FetchUserIdInBlacklist(t *testing.T) {
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
		BlacklistRepository repo.IBlacklistRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint
		wantErr bool
	}{
		{
			"good",
			fields{
				BlacklistRepository: repo.NewBlacklistRepository(db),
			},
			[]uint{2, 4, 6, 8},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlacklistDataService{
				BlacklistRepository: tt.fields.BlacklistRepository,
			}
			got, err := b.FetchUserIdInBlacklist()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUserIdInBlacklist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUserIdInBlacklist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlacklistDataService_IsUserIdInBlacklist(t *testing.T) {
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
		BlacklistRepository repo.IBlacklistRepository
	}
	type args struct {
		userId uint64
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
			fields{
				BlacklistRepository: repo.NewBlacklistRepository(db),
			},
			args{
				2,
			},
			true,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlacklistDataService{
				BlacklistRepository: tt.fields.BlacklistRepository,
			}
			got, err := b.IsUserIdInBlacklist(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserIdInBlacklist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsUserIdInBlacklist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestNewBlacklistDataService(t *testing.T) {
//	type args struct {
//		blacklistRepository repo.IBlacklistRepository
//	}
//	tests := []struct {
//		name string
//		args args
//		want IBlacklistDataService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewBlacklistDataService(tt.args.blacklistRepository); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewBlacklistDataService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
