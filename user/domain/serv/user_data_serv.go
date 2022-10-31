package serv

import (
	"context"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/cuno-1000/panic-product/user/domain/model"
	"github.com/cuno-1000/panic-product/user/domain/repo"
	pb "github.com/cuno-1000/panic-product/user/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

var RepaymentServiceConn *grpc.ClientConn

type IUserDataService interface {
	// AddUser 插入用户
	AddUser(*model.User) (uint, error)
	// DeleteUser 删除用户
	DeleteUser(uint) error
	// UpdateUser 更新用户
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	// FindUserById 根据id称查找用信息
	FindUserById(uint) (*model.User, error)
	// FindUserByTel 根据电话号码称查找用信息
	FindUserByTel(string) (*model.User, error)
	// FindUserByIdNumber 根据身份证号码称查找用户信息
	FindUserByIdNumber(string) (*model.User, error)
	// CheckPwd 比对账号密码是否正确
	CheckPwd(userTel string, pwd string) (isOk bool, id uint)
}

// NewUserDataService 创建实例
func NewUserDataService(userRepository repo.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repo.IUserRepository
}

// GeneratePassword 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

// AddUser 插入用户
func (u *UserDataService) AddUser(user *model.User) (userID uint, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}

	user.HashPassword = string(pwdByte)

	userID, err = u.UserRepository.CreateUser(user)

	client := pb.NewRepaymentRecordClient(RepaymentServiceConn)
	reply, _ := client.SeedRecord(context.Background(), &pb.SeedRecordRequest{
		UserId: uint64(userID),
	})
	if !reply.IsSuccess {
		fmt.Println("FAIL TO SEED REPAYMENT RECORD")
	}
	return
}

// DeleteUser 删除用户
func (u *UserDataService) DeleteUser(userID uint) error {
	return u.UserRepository.DeleteUserByID(userID)
}

// UpdateUser 更新用户
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

// FindUserByTel 根据用户名称查找用信息
func (u *UserDataService) FindUserByTel(userTel string) (user *model.User, err error) {
	return u.UserRepository.FindUserByTel(userTel)
}

// FindUserById 根据用户名称查找用信息
func (u *UserDataService) FindUserById(userId uint) (user *model.User, err error) {
	return u.UserRepository.FindUserByID(userId)
}

// FindUserByIdNumber 根据身份证号码称查找用户信息
func (u *UserDataService) FindUserByIdNumber(userIdNo string) (user *model.User, err error) {
	return u.UserRepository.FindUserByTel(userIdNo)
}

// CheckPwd 比对账号密码是否正确
func (u *UserDataService) CheckPwd(userTel string, pwd string) (isOk bool, id uint) {
	user, err := u.UserRepository.FindUserByTel(userTel)
	if err != nil {
		return false, 0
	}
	isOk, err = ValidatePassword(pwd, user.HashPassword)
	logger.Error(err)
	return isOk, user.ID
}
