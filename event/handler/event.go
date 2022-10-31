package handler

import (
	"context"
	"github.com/cuno-1000/panic-product/event/domain/model"
	"github.com/cuno-1000/panic-product/event/domain/serv"
	event "github.com/cuno-1000/panic-product/event/proto"
	"github.com/shopspring/decimal"
	"time"
)

type Event struct {
	EventDataService serv.IEventDataService
}

func (e *Event) CreateEvent(_ context.Context, in *event.CreateEventRequest) (out *event.CreateEventResponse, _ error) {
	out = &event.CreateEventResponse{}
	timeStamp1, err := time.Parse("2006-01-02 15:04:05", in.ReviewUpperLimitAt)
	timeStamp2, err := time.Parse("2006-01-02 15:04:05", in.StartingAt)
	if err != nil {
		return out, err
	}
	eventSetUp := &model.Event{
		ID:                          uint(in.Id),
		Info:                        in.Info,
		AdminId:                     uint(in.AdminId),
		ApplyRules:                  in.ApplyRules,
		ProductQuantity:             in.ProductQuantity,
		ProductItemPrice:            decimal.NewFromFloatWithExponent(in.ProductItemPrice, 0),
		RepaymentReviewUpperLimitAt: timeStamp1,
		StartingAt:                  timeStamp2,
		OverDueMaxTimes:             uint(in.OverDueMaxTimes),
	}
	out.Uuid, out.IsSuccess = e.EventDataService.CreateEvent(eventSetUp)
	return out, nil
}

func (e *Event) FetchEvents(_ context.Context, _ *event.FetchEventsRequest) (out *event.FetchEventsResponse, _ error) {
	out = &event.FetchEventsResponse{}
	var err error
	out.Events, err = e.EventDataService.FetchEvents()
	if err != nil {
		return out, err
	}
	return out, nil
}

func (e *Event) CheckBlacklist(_ context.Context, in *event.CheckBlacklistRequest) (out *event.CheckBlacklistResponse, _ error) {
	out = &event.CheckBlacklistResponse{}
	eventModel := e.EventDataService.EventByUuid(in.Uuid)

	temp, err := e.EventDataService.CheckBlacklist(eventModel, in.UserId)
	if err != nil {
		out.Situation = int64(temp)
		return out, err
	}
	out.Situation = int64(temp)
	var m = serv.RejectedSituationMap()

	if out.Situation == int64(m["success"]) {
		out.Link, _ = e.EventDataService.CreateApplyLink(eventModel, in.UserId)
		return
	}
	return out, nil
}

func (e *Event) CheckBlacklistDowngrade(_ context.Context, in *event.CheckBlacklistRequest) (out *event.CheckBlacklistResponse, _ error) {
	out = &event.CheckBlacklistResponse{}
	eventModel := e.EventDataService.EventByUuid(in.Uuid)

	temp, err := e.EventDataService.CheckBlacklist(eventModel, in.UserId)
	if err != nil {
		out.Situation = int64(temp)
		return out, err
	}
	out.Situation = int64(temp)
	var m = serv.RejectedSituationMap()

	if out.Situation == int64(m["success"]) {
		out.Link, _ = e.EventDataService.CreateApplyLink(eventModel, in.UserId)
		return
	}
	return out, nil
}

func (e *Event) ApplyPurchase(_ context.Context, in *event.ApplyRequest) (out *event.ApplyResponse, _ error) {
	out = &event.ApplyResponse{}
	m := serv.RejectedSituationMap()

	eventModel := e.EventDataService.EventByUuid(in.Uuid)
	if in.Link == "" {
		_, err := e.EventDataService.CheckBlacklist(eventModel, in.UserId)
		if err != nil {
			return nil, err
		}
		//if situation > -1 {
		//	//var eventApply *model.EventApplyRecord
		//	//eventApply = &model.EventApplyRecord{
		//	//	EventId:   uint64(eventModel.ID),
		//	//	UserId:    in.UserId,
		//	//	Situation: uint64(situation),
		//	//}
		//	//e.EventDataService.PushToMq(eventApply)
		//}

	} else if !e.EventDataService.CheckApplyLink(in.Uuid, in.UserId, in.Link) {
		out.Situation = int64(m["not_start_yet"])
		return
	}

	now := time.Now()
	if eventModel.StartingAt.IsZero() || now.Before(eventModel.StartingAt) {
		out.Situation = int64(m["not_start_yet"])
		return out, nil
	}

	out.Situation = e.EventDataService.DecreaseStock(eventModel, in.UserId)

	if out.Situation == int64(m["success"]) {
		out.IsSuccess = true
	} else {
		out.IsSuccess = false
	}

	if out.Situation < int64(m["limit_purchase"]) {
		var eventApply *model.EventApplyRecord
		eventApply = &model.EventApplyRecord{
			EventId:   uint64(eventModel.ID),
			UserId:    in.UserId,
			Situation: out.Situation,
		}
		e.EventDataService.PushToMq(eventApply)
	}

	return out, nil
}
