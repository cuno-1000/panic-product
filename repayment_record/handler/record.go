package handler

import (
	"context"
	"github.com/cuno-1000/panic-product/repayment_record/domain/serv"
	repayment "github.com/cuno-1000/panic-product/repayment_record/proto"
	"math/rand"
	"time"
)

type Record struct {
	RecordDataService    serv.IRecordDataService
	BlacklistDataService serv.IBlacklistDataService
}

func (r *Record) SeedRecord(_ context.Context, in *repayment.SeedRecordRequest) (out *repayment.SeedRecordResponse, _ error) {
	out = &repayment.SeedRecordResponse{}
	if rand.Intn(10) == 1 {
		_ = r.BlacklistDataService.AppendBlacklist(uint(in.UserId))
	}
	if err := r.RecordDataService.AddRecord(uint(in.UserId)); err != nil {
		out.IsSuccess = false
		return out, nil
	}
	out.IsSuccess = true
	return
}

// FetchBlacklist 获得失信被执行人名单中所有用户
func (r *Record) FetchBlacklist(_ context.Context, _ *repayment.BlacklistRequest) (out *repayment.BlacklistResponse, _ error) {
	blacklistUserIdAll, err := r.BlacklistDataService.FetchUserIdInBlacklist()
	out = &repayment.BlacklistResponse{}
	if err != nil {
		return out, err
	}
	for _, v := range blacklistUserIdAll {
		out.UserId = append(out.UserId, uint64(v))
	}
	return out, err
}

// FetchRepaymentOverDue 获取还款记录中需被筛除的用户
func (r *Record) FetchRepaymentOverDue(_ context.Context, in *repayment.RepaymentOverDueRequest) (out *repayment.RepaymentOverDueResponse, _ error) {
	out = &repayment.RepaymentOverDueResponse{}
	timeUpperLimit, err := time.Parse("2006-01-02 15:04:05", in.TimeUpperLimit)
	if err != nil {
		return nil, err
	}
	overDueUserIdAll, err := r.RecordDataService.GetUserIDNotAcceptedToPanicEvent(timeUpperLimit, in.AdaptedRemainingRepayments, uint(in.MaxTimes))
	if err != nil {
		return out, err
	}
	for _, v := range overDueUserIdAll {
		out.UserId = append(out.UserId, uint64(v))
	}
	return out, err
}

// IsUserIdRepaymentOverDue 判断还款记录中用户是否需被筛除
func (r *Record) IsUserIdRepaymentOverDue(_ context.Context, in *repayment.IsUserIdRepaymentOverDueRequest) (out *repayment.IsUserIdRepaymentOverDueResponse, _ error) {
	out = &repayment.IsUserIdRepaymentOverDueResponse{}
	var err error
	timeUpperLimit, err := time.Parse("2006-01-02 15:04:05", in.TimeUpperLimit)
	if err != nil {
		return nil, err
	}
	out.OverDue, err = r.RecordDataService.IsAcceptedToPanicEventByUserId(timeUpperLimit, in.UserId, in.AdaptedRemainingRepayments, uint(in.MaxTimes))
	return out, err
}

// IsInBlacklist 判断用户是否在失信被执行人名单中
func (r *Record) IsInBlacklist(_ context.Context, in *repayment.IsInBlacklistRequest) (out *repayment.IsInBlacklistResponse, _ error) {
	out = &repayment.IsInBlacklistResponse{}
	var err error
	out.In, err = r.BlacklistDataService.IsUserIdInBlacklist(in.UserId)
	return out, err
}
