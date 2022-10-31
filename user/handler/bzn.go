package handler

import (
	"context"
	"fmt"
	user "github.com/cuno-1000/panic-product/user/proto"
)

// CheckInfo 身份证获取信息
func (u *User) CheckInfo(_ context.Context, in *user.UserCheckInfoRequest) (out *user.UserCheckInfoResponse, _ error) {
	out = &user.UserCheckInfoResponse{}
	userModel, err := u.UserDataService.FindUserById(uint(in.UserId))
	if err != nil {
		return out, nil
	}
	IdNumber := userModel.IdNumber

	age := u.BznDataService.AgeCounter(IdNumber)
	out.Age = uint32(age)
	out.IsAdult = age > 17
	out.Gender = u.BznDataService.GetGender(IdNumber)
	out.IsCareerStatusNormal = userModel.CareerStatus == ""
	return out, nil
}

// FetchNormalAdult 获取此刻所有用户
func (u *User) FetchNormalAdult(_ context.Context, _ *user.FetchNormalAdultRequest) (out *user.FetchNormalAdultResponse, _ error) {
	out = &user.FetchNormalAdultResponse{}
	var careerSituation, notAdultSituation int64
	careerSituation = 16
	notAdultSituation = 8

	users, _ := u.BznDataService.FetchNormalAdult()
	for _, v := range users {
		normalAdultItem := &user.UserWithAgeCareerSituation{}
		if u.BznDataService.AgeCounter(v.IdNumber) < 18 {
			normalAdultItem.Situation += notAdultSituation
		}
		normalAdultItem.UserId = uint64(v.ID)
		if v.CareerStatus != "" {
			normalAdultItem.Situation += careerSituation
		}
		out.UserWithSituation = append(out.UserWithSituation, normalAdultItem)
	}
	return
}

// ReduceBalance 减少余额
func (u *User) ReduceBalance(_ context.Context, in *user.ReduceBalanceRequest) (out *user.ReduceBalanceResponse, _ error) {
	out = &user.ReduceBalanceResponse{}
	var err error
	out.IsSuccess, err = u.BznDataService.ReduceUserBalance(uint(in.UserId), in.Amount)
	if err != nil {
		fmt.Println(err)
		return out, err
	}
	return out, nil
}
