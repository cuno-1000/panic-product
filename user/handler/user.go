package handler

import (
	"context"
	"github.com/cuno-1000/panic-product/user/domain/model"
	"github.com/cuno-1000/panic-product/user/domain/serv"
	user "github.com/cuno-1000/panic-product/user/proto"
)

type User struct {
	UserDataService serv.IUserDataService
	BznDataService  serv.IBznDataService
}

// Register 注册
func (u *User) Register(_ context.Context, in *user.UserRegisterRequest) (out *user.UserRegisterResponse, _ error) {
	userRegister := &model.User{
		Tel:          in.UserTel,
		Name:         in.UserName,
		IdNumber:     in.UserIdNumber,
		HashPassword: in.UserPwd,
	}
	out = &user.UserRegisterResponse{}
	if !u.BznDataService.ValidateIdCartNumber(userRegister.IdNumber) {
		out.Message = "身份证无效"
		return out, nil
	} else {
		userRegister.IdNumber = userRegister.IdNumber[:17]
	}
	id, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		out.Message = "电话号码已存在"
		return out, nil
	}
	out.Message = "success"
	out.UserId = uint64(id)
	return out, nil
}

// Login 登录
func (u *User) Login(_ context.Context, in *user.UserLoginRequest) (out *user.UserLoginResponse, _ error) {
	tel := in.UserTel
	pwd := in.Pwd
	isOk, id := u.UserDataService.CheckPwd(tel, pwd)
	out = &user.UserLoginResponse{}
	out.IsSuccess = isOk
	out.UserId = uint64(id)
	return out, nil
}

// GetUserInfo 查询用户信息
func (u *User) GetUserInfo(_ context.Context, in *user.UserInfoRequest) (out *user.UserInfoResponse, _ error) {
	userInfo, err := u.UserDataService.FindUserById(uint(in.UserId))
	out = &user.UserInfoResponse{}
	if err != nil {
		return out, err
	}
	UserForResponse(userInfo, out)
	return out, nil
}

// UserForResponse 响应内容
func UserForResponse(userModel *model.User, response *user.UserInfoResponse) {
	response.UserTel = userModel.Tel
	response.UserName = userModel.Name
	response.UserIdNumber = userModel.IdNumber + serv.GetIdNumberValidateCODE(userModel.IdNumber)
	response.UserCareerStatus = userModel.CareerStatus
	response.UserBalance = userModel.AccountBalance.String()
}

func (u *User) TestService(_ context.Context, in *user.TestingRequest) (out *user.TestingResponse, _ error) {
	out = &user.TestingResponse{}
	if in.IsSuccess {
		out.Message = "It is wonderful."
		return out, nil
	} else {
		return nil, nil
	}
}
