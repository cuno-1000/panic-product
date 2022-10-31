package serv

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/cuno-1000/panic-product/event/domain/model"
	"github.com/cuno-1000/panic-product/event/domain/repo"
	pb "github.com/cuno-1000/panic-product/event/proto"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

type IEventDataService interface {
	CreateEvent(event *model.Event) (string, bool)
	WarmUpEventRoutines()
	EventByUuid(uuid string) *model.Event
	FetchEvents() ([]*pb.EventItem, error)
	WarmUpBlacklist(*model.Event) (bool, error)
	WarmUpRepaymentOverDue(event *model.Event) (bool, error)
	AppendUserToBlacklist(event *model.Event, userId uint64, situation int) (isDone bool)
	CheckBlacklist(event *model.Event, userId uint64) (int, error)
	CheckBlacklistDowngrade(event *model.Event, userId uint64) (situation int, err error)
	CreateApplyLink(event *model.Event, userId uint64) (string, error)
	CheckApplyLink(uuid string, userId uint64, linkModel string) bool
	DecreaseStock(event *model.Event, userId uint64) int64
	PushToMq(record *model.EventApplyRecord) bool
}

func NewEventDataService(eventRepository repo.IEventRepository) IEventDataService {
	return &EventDataService{EventRepository: eventRepository}
}

type EventDataService struct {
	EventRepository repo.IEventRepository
}

func RejectedSituationMap() map[string]int {
	return map[string]int{
		"success":              1,
		"repayment":            2,
		"blacklist":            4,
		"age":                  8,
		"career":               16,
		"limit_purchase":       32,
		"not_start_yet":        64,
		"insufficient_balance": 128,
		"sold_out":             256,
	}
}

// CreateEvent 创建活动缓存
func (e *EventDataService) CreateEvent(event *model.Event) (string, bool) {
	event.Uuid = CreateUUID()
	isSuccess, err := e.EventRepository.CreateEvent(event)
	if err != nil {
		return "", false
	}
	if isSuccess {
		NewEventWarmUpTask <- event.Uuid
		return "正在预热用户数据", true
	} else {
		return "预热用户数据失败", true
	}
}

// EventByUuid 通过UUID获取活动缓存信息
func (e *EventDataService) EventByUuid(uuid string) (event *model.Event) {
	event = &model.Event{}
	event, _ = e.EventRepository.FindEventByUuid(uuid)
	return event
}

// FetchEvents 获取活动缓存列表
func (e *EventDataService) FetchEvents() (el []*pb.EventItem, err error) {
	var elItem pb.EventItem
	events, err := e.EventRepository.FetchEvents()
	if err != nil {
		return nil, err
	}
	for _, v := range events {
		elItem.UuidUrl = v.Uuid
		elItem.Info = v.Info
		el = append(el, &elItem)
	}
	return
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

// CheckBlacklist 获取某用户在特定活动的准入情况
func (e *EventDataService) CheckBlacklist(event *model.Event, userId uint64) (situation int, err error) {
	m := RejectedSituationMap()
	if situation, err = e.EventRepository.FindInBlacklist("EVENT_"+event.Uuid+"_blacklists", userId); err == redis.Nil {

		client1 := pb.NewUserClient(UserServiceConn)
		if step1, err := client1.CheckInfo(context.Background(), &pb.UserCheckInfoRequest{
			UserId: userId,
		}); err != nil {
			return -1, err
		} else {
			if !step1.IsAdult {
				situation += m["age"]
			}
			if !step1.IsCareerStatusNormal {
				situation += m["career"]
			}
		}

		client2 := pb.NewRepaymentRecordClient(RepaymentServiceConn)
		if step3, err := client2.IsInBlacklist(context.Background(), &pb.IsInBlacklistRequest{
			UserId: userId,
		}); err != nil {
			return -1, err
		} else {
			if step3.In {
				situation += m["blacklist"]
			}
		}

		if step3, err := client2.IsUserIdRepaymentOverDue(context.Background(), &pb.IsUserIdRepaymentOverDueRequest{
			TimeUpperLimit:             event.RepaymentReviewUpperLimitAt.Format("2006-01-02 15:04:05"),
			AdaptedRemainingRepayments: event.ApplyRules,
			UserId:                     userId,
			MaxTimes:                   uint32(event.OverDueMaxTimes),
		}); err != nil {
			return -1, err
		} else {
			if step3.OverDue {
				situation += m["repayment"]
			}
		}

		e.AppendUserToBlacklist(event, userId, situation)

	} else if err != nil {
		situation = -1
	}
	if situation == 0 {
		situation = m["success"]
	}
	return situation, nil
}

// CheckBlacklistDowngrade 获取某用户在特定活动的准入情况
func (e *EventDataService) CheckBlacklistDowngrade(event *model.Event, userId uint64) (situation int, err error) {
	m := RejectedSituationMap()
	if situation, err = e.EventRepository.FindInBlacklist("EVENT_"+event.Uuid+"_blacklists", userId); err == redis.Nil {

		client1 := pb.NewUserClient(UserServiceConn)
		if step1, err := client1.CheckInfo(context.Background(), &pb.UserCheckInfoRequest{
			UserId: userId,
		}); err != nil {
			return -1, err
		} else {
			if !step1.IsAdult {
				situation += m["age"]
			}
			if !step1.IsCareerStatusNormal {
				situation += m["career"]
			}
		}
		if situation != 0 {
			e.AppendUserToBlacklist(event, userId, situation)
			return situation, nil
		}

		client2 := pb.NewRepaymentRecordClient(RepaymentServiceConn)
		if step3, err := client2.IsInBlacklist(context.Background(), &pb.IsInBlacklistRequest{
			UserId: userId,
		}); err != nil {
			return -1, err
		} else {
			if step3.In {
				situation += m["blacklist"]
			}
		}
		if situation != 0 {
			e.AppendUserToBlacklist(event, userId, situation)
			return situation, nil
		}

		if step3, err := client2.IsUserIdRepaymentOverDue(context.Background(), &pb.IsUserIdRepaymentOverDueRequest{
			TimeUpperLimit:             event.RepaymentReviewUpperLimitAt.Format("2006-01-02 15:04:05"),
			AdaptedRemainingRepayments: event.ApplyRules,
			UserId:                     userId,
			MaxTimes:                   uint32(event.OverDueMaxTimes),
		}); err != nil {
			return -1, err
		} else {
			if step3.OverDue {
				situation += m["repayment"]
			}
		}
		if situation != 0 {
			e.AppendUserToBlacklist(event, userId, situation)
			return situation, nil
		}

		e.AppendUserToBlacklist(event, userId, situation)

	} else if err != nil {
		situation = -1
	}
	if situation == 0 {
		situation = m["success"]
	}
	return situation, nil
}

func (e *EventDataService) CreateApplyLink(event *model.Event, userId uint64) (string, error) {
	if link, err := e.EventRepository.GetLink("EVENT_"+event.Uuid+"_links", userId); err != redis.Nil {
		return link, nil
	} else {
		link = CreateUUID()
		if isOk, _ := e.EventRepository.SetLink("EVENT_"+event.Uuid+"_links", userId, link); isOk {
			return link, nil
		}
		return "", nil
	}
}

func (e *EventDataService) CheckApplyLink(uuid string, userId uint64, linkModel string) bool {
	if link, err := e.EventRepository.GetLink("EVENT_"+uuid+"_links", userId); err == redis.Nil {
		return false
	} else if link == linkModel {
		return true
	}
	return false
}

//var grandChannel = make(chan struct{}, 1)

//var mutex sync.Mutex

func getQuantityWithLock(e *EventDataService, event *model.Event, locker *sync.Mutex) (isFull bool) {
	locker.Lock()
	defer locker.Unlock()
	if quantity, _ := e.EventRepository.OrdersQuantity("EVENT_"+event.Uuid+"_orders", int64(event.ProductQuantity)); uint32(quantity) == event.ProductQuantity {
		isFull = true
	}
	return
}

func (e *EventDataService) DecreaseStock(event *model.Event, userId uint64) int64 {

	//var quantity int64
	m := RejectedSituationMap()

	//mutex.Lock()
	////grandChannel <- struct{}{}
	//
	//quantity, _ := e.EventRepository.OrdersQuantity("EVENT_"+event.Uuid+"_orders", int64(event.ProductQuantity))
	//if uint32(quantity) == event.ProductQuantity {
	//	mutex.Unlock()
	//	//<-grandChannel
	//	return int64(m["sold_out"])
	//}
	////mutex.Unlock()
	////<-grandChannel
	////
	////if getQuantityWithLock(e, event, &mutex) {
	////	return int64(m["sold_out"])
	////}
	//
	situation, _ := e.EventRepository.DecreaseStork(event.Uuid, userId)
	//<-grandChannel
	//mutex.Unlock()

	if situation == 1 {
		client := pb.NewUserClient(UserServiceConn)
		rsp, err := client.ReduceBalance(context.Background(), &pb.ReduceBalanceRequest{
			UserId: userId,
			Amount: event.ProductItemPrice.String(),
		})
		if err != nil {
			return -1
		}
		if !rsp.IsSuccess {
			e.EventRepository.RollbackStork("EVENT_"+event.Uuid+"_orders", userId)
			return int64(m["insufficient_balance"])
		}
	} else if situation == 0 {
		return int64(m["limit_purchase"])
	} else if situation == 2 {
		return int64(m["sold_out"])
	}
	//event.ProductItemPrice.Mul(decimal.New(-1, 0)).String()
	return int64(m["success"])
}
