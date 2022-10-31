package handler

import (
	"context"
	user "github.com/cuno-1000/panic-product/api-gateway/proto"
	"github.com/cuno-1000/panic-product/api-gateway/token/middleware"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"

	//"github.com/cuno-1000/panic-product/api-gateway/token/common/request"
	//"github.com/cuno-1000/panic-product/api-gateway/token/common/response"
	"github.com/cuno-1000/panic-product/api-gateway/token/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var UserServiceConn *grpc.ClientConn

type LoginForm struct {
	Tel string `json:"tel"`
	Pwd string `json:"pwd"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.BindJSON(&form); err != nil {
		Fail(c, "Cannot decode request")
		return
	}

	client := user.NewUserClient(UserServiceConn)
	if reply, err := client.Login(context.Background(), &user.UserLoginRequest{
		UserTel: form.Tel,
		Pwd:     form.Pwd,
	}); err != nil || !reply.IsSuccess {
		Fail(c, "用户不存在或密码错误")
		return
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, reply.UserId)
		if err != nil {
			Fail(c, "Cannot create token")
			return
		}
		Success(c, tokenData)
	}
}

type RegisterForm struct {
	UserName     string `json:"name"`
	UserTel      string `json:"tel"`
	UserIdNumber string `json:"user_id_number"`
	Pwd          string `json:"pwd"`
}

func Register(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		Fail(c, "Cannot decode request")
		return
	}

	client := user.NewUserClient(UserServiceConn)
	if reply, err := client.Register(context.Background(), &user.UserRegisterRequest{
		UserTel:      form.UserTel,
		UserName:     form.UserName,
		UserIdNumber: form.UserIdNumber,
		UserPwd:      form.Pwd,
	}); err != nil {
		Fail(c, err)
		return
	} else if reply.Message != "success" {
		Fail(c, reply.Message)
		return
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, reply.UserId)
		if err != nil {
			Fail(c, "Cannot create token")
			return
		}
		Success(c, tokenData)
	}
}

func GetUserInfo(c *gin.Context) {
	id, err := middleware.IdFromJWT(c)
	if err != nil {
		Fail(c, "Cannot read token")
		c.Abort()
		return
	}

	client := user.NewUserClient(UserServiceConn)
	rsp, err := client.GetUserInfo(context.Background(), &user.UserInfoRequest{
		UserId: id,
	})
	if err != nil {
		Fail(c, err)
		return
	}
	Success(c, rsp)
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Log out failed")
		return
	}
}

func Test(c *gin.Context) {
	client := user.NewUserClient(UserServiceConn)
	if rsp, err := client.TestService(context.Background(), &user.TestingRequest{
		IsSuccess: true,
	}); err != nil {
		Fail(c, err)
		return
	} else {
		Success(c, rsp.Message)
	}
}

func FetchNormalAdult(c *gin.Context) {
	client := user.NewUserClient(UserServiceConn)
	if rsp, err := client.FetchNormalAdult(context.Background(), &user.FetchNormalAdultRequest{}); err != nil {
		Fail(c, err)
		return
	} else {
		Success(c, rsp.UserWithSituation)
	}
}
