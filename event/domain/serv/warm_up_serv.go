package serv

import (
	"context"
	"fmt"
	"github.com/cuno-1000/panic-product/event/domain/model"
	pb "github.com/cuno-1000/panic-product/event/proto"
	"google.golang.org/grpc"
)

var UserServiceConn *grpc.ClientConn
var RepaymentServiceConn *grpc.ClientConn

var NewEventWarmUpTask = make(chan string)

func (e *EventDataService) WarmUpEventRoutines() {
	var eventUuid string
	for {
		select {
		case eventUuid = <-NewEventWarmUpTask:
			eventModel := e.EventByUuid(eventUuid)
			///////////
			e.WarmUpNormalAdult(eventModel)
			e.WarmUpBlacklist(eventModel)
			e.WarmUpRepaymentOverDue(eventModel)
		}
	}
}

func (e *EventDataService) WarmUpNormalAdult(event *model.Event) (isDone bool, _ error) {
	client := pb.NewUserClient(UserServiceConn)
	if reply, err := client.FetchNormalAdult(context.Background(), &pb.FetchNormalAdultRequest{}); err != nil {
		return false, err
	} else {
		for _, v := range reply.UserWithSituation {
			isDone = e.EventRepository.CreateBlacklist("EVENT_"+event.Uuid+"_blacklists", v.UserId, v.Situation)
		}
	}
	return true, nil
}

// WarmUpBlacklist 预热失信被执行人名单到活动缓存
func (e *EventDataService) WarmUpBlacklist(event *model.Event) (isDone bool, _ error) {

	client := pb.NewRepaymentRecordClient(RepaymentServiceConn)
	if reply, err := client.FetchBlacklist(context.Background(), &pb.BlacklistRequest{}); err != nil {
		return false, err
	} else {
		m := RejectedSituationMap()
		isDone = e.EventRepository.PlusAppendBlacklist("EVENT_"+event.Uuid+"_blacklists", reply.UserId, int64(m["blacklist"]))
		if err != nil {
			return false, err
		}
		if isDone {
			fmt.Println("Blacklist Warmed Up")
		}
	}
	return true, nil
}

// WarmUpRepaymentOverDue 预热还款记录中需被筛除的用户到活动缓存
func (e *EventDataService) WarmUpRepaymentOverDue(event *model.Event) (isDone bool, _ error) {
	client := pb.NewRepaymentRecordClient(RepaymentServiceConn)
	if reply, err := client.FetchRepaymentOverDue(context.Background(), &pb.RepaymentOverDueRequest{
		TimeUpperLimit:             event.RepaymentReviewUpperLimitAt.Format("2006-01-02 15:04:05"),
		AdaptedRemainingRepayments: event.ApplyRules,
		MaxTimes:                   uint32(event.OverDueMaxTimes),
	}); err != nil {
		return false, err
	} else {
		m := RejectedSituationMap()
		isDone = e.EventRepository.PlusAppendBlacklist("EVENT_"+event.Uuid+"_blacklists", reply.UserId, int64(m["repayment"]))
		if err != nil {
			return false, err
		}
		if isDone {
			fmt.Println("Blacklist Warmed Up")
		}
	}
	return true, nil
}

func (e *EventDataService) AppendUserToBlacklist(event *model.Event, userId uint64, situation int) (isDone bool) {
	return e.EventRepository.AppendElement("EVENT_"+event.Uuid+"_blacklists", userId, int64(situation))
}
