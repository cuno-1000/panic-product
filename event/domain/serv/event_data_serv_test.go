package serv

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/cuno-1000/panic-product/event/domain/model"
	"github.com/cuno-1000/panic-product/event/domain/repo"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"reflect"
	"strconv"
	"testing"
	"time"
)

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

var RedisDb = redisInit()

func TestEventDataService_CreateEvent1(t *testing.T) {
	d, _ := time.ParseDuration("-24h")
	var event = &model.Event{
		ID:                          1,
		Info:                        "BBBBBBBBB",
		AdminId:                     500,
		ApplyRules:                  "date_difference < 3 AND owed_amount < 3000",
		ProductQuantity:             10000,
		ProductItemPrice:            decimal.NewFromFloatWithExponent(1000, 0),
		RepaymentReviewUpperLimitAt: time.Now().Add(time.Duration(1) * d),
		OverDueMaxTimes:             2,
		StartingAt:                  time.Now(),
	}
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		event *model.Event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			"good",
			fields{
				EventRepository: repo.NewEventRepository(RedisDb),
			},
			args{
				event,
			},
			"OK",
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			got, got1 := e.CreateEvent(tt.args.event)
			if got != tt.want {
				t.Errorf("CreateEvent() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateEvent() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEventDataService_EventByUuid(t *testing.T) {
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Event
	}{
		{
			"good",
			fields{
				EventRepository: repo.NewEventRepository(RedisDb),
			},
			args{
				"94bde9c6-3b8a-4a7d-6a6b-5a5f9a59c9f3",
			},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			if got := e.EventByUuid(tt.args.uuid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventByUuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventDataService_CheckApplyLink(t *testing.T) {
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		uuid      string
		userId    uint64
		linkModel string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"good",
			fields{
				EventRepository: repo.NewEventRepository(RedisDb),
			},
			args{
				"8464da55-5be3-4acd-7bc2-90dba0af2769",
				11,
				"0a264e9f-45b6-462c-68df-fd4bf7ec5934",
			},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			if got := e.CheckApplyLink(tt.args.uuid, tt.args.userId, tt.args.linkModel); got != tt.want {
				t.Errorf("CheckApplyLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
