package repo

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/cuno-1000/panic-product/event/domain/model"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"log"
	"strconv"
	"time"
)

type IEventRepository interface {
	CreateEvent(event *model.Event) (bool, error)
	FindEventByUuid(uuid string) (*model.Event, error)
	FetchEvents() ([]model.Event, error)
	CreateBlacklist(prefix string, userId uint64, situation int64) bool
	PlusAppendBlacklist(prefix string, userId []uint64, situation int64) bool
	AppendElement(prefix string, userId uint64, situation int64) bool

	FindInBlacklist(prefix string, userId uint64) (situation int, err error)

	SetLink(prefix string, userId uint64, link string) (bool, error)
	GetLink(prefix string, userId uint64) (link string, err error)

	OrdersQuantity(prefix string, limit int64) (int64, error)

	IsOrderedUserId(prefix string, userId uint64) (is bool, err error)

	DecreaseStork(prefix string, userId uint64) (int, error)

	RollbackStork(prefix string, userId uint64) (isSuccess bool, err error)
}

func NewEventRepository(db *redis.Client) IEventRepository {
	return &EventRepository{redisDb: db}
}

type EventRepository struct {
	redisDb *redis.Client
}

func (e *EventRepository) CreateEvent(event *model.Event) (bool, error) {
	var err error
	event.Uuid = CreateUUID()
	var key = "EVENT_INFO_" + event.Uuid
	if err = e.redisDb.HSet(context.Background(), key,
		"id", event.ID,
		"info", event.Info,
		"admin_id", event.AdminId,
		"apply_rules", event.ApplyRules,
		"product_quantity", event.ProductQuantity,
		"product_item_price", event.ProductItemPrice.String(),
		//"created_at", event.CreatedAt,
		"repayment_search_upper_time", event.RepaymentReviewUpperLimitAt.Format("2006-01-02 15:04:05"),
		"over_due_max_times", event.OverDueMaxTimes,
		"starting_at", event.StartingAt.Format("2006-01-02 15:04:05"),
	).Err(); err != nil {
		return false, err
	}
	return true, nil
}

// CreateUUID create a random UUID with from RFC 4122
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func (e *EventRepository) FindEventByUuid(uuid string) (*model.Event, error) {
	var event = &model.Event{}
	var key = "EVENT_INFO_" + uuid
	result, err := e.redisDb.HGetAll(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, err
	}
	event.StartingAt, _ = time.Parse("2006-01-02 15:04:05", result["starting_at"])
	if event.StartingAt.IsZero() {
		return nil, errors.New("Internal Error. ")
	}
	event.Uuid = uuid
	event.Info = result["info"]
	temp, _ := strconv.ParseUint(result["admin_id"], 10, 32)
	event.AdminId = uint(temp)
	event.ApplyRules = result["apply_rules"]
	temp2, _ := strconv.ParseUint(result["product_quantity"], 10, 32)
	event.ProductQuantity = uint32(temp2)
	event.ProductItemPrice, _ = decimal.NewFromString(result["product_item_price"])
	event.RepaymentReviewUpperLimitAt, _ = time.Parse("2006-01-02 15:04:05", result["repayment_search_upper_time"])
	temp3, _ := strconv.ParseUint(result["over_due_max_times"], 10, 32)
	event.OverDueMaxTimes = uint(temp3)
	temp4, _ := strconv.ParseUint(result["id"], 10, 32)
	event.ID = uint(temp4)
	return event, nil
}

func (e *EventRepository) FetchEvents() (events []model.Event, err error) {
	var uuidAll []string
	var temp string
	var event model.Event
	uuidAll, err = e.redisDb.Keys(context.Background(), "EVENT_INFO_*").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	for _, v := range uuidAll {
		temp, err = e.redisDb.HGet(context.Background(), v, "info").Result()
		event = model.Event{
			Uuid: v[11:],
			Info: temp,
		}
		events = append(events, event)
	}
	return
}

func (e *EventRepository) CreateBlacklist(prefix string, userId uint64, situation int64) bool {
	str, _ := e.redisDb.HSet(context.Background(), prefix, userId, situation).Result()
	if str == 1 {
		return true
	}
	return false
}

func (e *EventRepository) PlusAppendBlacklist(prefix string, userId []uint64, situation int64) bool {
	for _, v := range userId {
		if isExists, _ := e.redisDb.HExists(context.Background(), prefix, strconv.FormatUint(v, 10)).Result(); isExists {
			e.redisDb.HIncrBy(context.Background(), prefix, strconv.FormatUint(v, 10), situation)
		}
	}
	return true
}

func (e *EventRepository) AppendElement(prefix string, userId uint64, situation int64) bool {
	num, _ := e.redisDb.HIncrBy(context.Background(), prefix, strconv.FormatUint(userId, 10), situation).Result()
	return num > -1
}

func (e *EventRepository) FindInBlacklist(prefix string, userId uint64) (situation int, err error) {
	str, err := e.redisDb.HGet(context.Background(), prefix, strconv.FormatUint(userId, 10)).Result()
	if err == redis.Nil {
		return 0, err
	}
	situation, err = strconv.Atoi(str)
	return situation, err
}

func (e *EventRepository) SetLink(prefix string, userId uint64, link string) (bool, error) {
	if str, err := e.redisDb.HSet(context.Background(), prefix, userId, link).Result(); err != nil {
		return false, nil
	} else if str == 1 {
		return true, nil
	}
	return false, nil
}

func (e *EventRepository) GetLink(prefix string, userId uint64) (link string, err error) {
	return e.redisDb.HGet(context.Background(), prefix, strconv.FormatUint(userId, 10)).Result()
}

func (e *EventRepository) OrdersQuantity(prefix string, limit int64) (int64, error) {
	return e.redisDb.SCard(context.Background(), prefix).Result()
}

func (e *EventRepository) IsOrderedUserId(prefix string, userId uint64) (is bool, err error) {
	is, _ = e.redisDb.SIsMember(context.Background(), prefix, userId).Result()
	return is, nil
}

func (e *EventRepository) DecreaseStork(uuid string, userId uint64) (num int, err error) {
	var decrStork = redis.NewScript(`
		local hkey = KEYS[1]
		local hfield = ARGV[1]
		local orders = KEYS[2]
		local elem = ARGV[2]
		
		local quantity = redis.call("SCARD", orders)
		local total = redis.call("HGET", hkey, hfield)
		if tonumber(quantity) == tonumber(total) then
			return "2"
		else
			return redis.call("SADD", orders, elem)
		end
	`)
	keys := []string{"EVENT_INFO_" + uuid, "EVENT_" + uuid + "_orders"}
	values := []interface{}{"product_quantity", userId}
	num, err = decrStork.Run(context.Background(), e.redisDb, keys, values...).Int()
	if err != nil {
		fmt.Println(num)
	}
	return num, nil
	//str, _ := e.redisDb.SAdd(context.Background(), prefix, userId).Result()
	//if str == 1 {
	//	return true, nil
	//}
	//return false, nil
}

func (e *EventRepository) RollbackStork(prefix string, userId uint64) (isSuccess bool, err error) {
	str, _ := e.redisDb.SRem(context.Background(), prefix, userId).Result()
	if str == 1 {
		return true, nil
	}
	return false, nil
}
