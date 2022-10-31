package repo

import (
	"github.com/cuno-1000/panic-product/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindUserByID 根据用户ID查找用户信息
	FindUserByID(uint) (*model.User, error)
	// FindUserByTel 根据电话号码称查找用户信息
	FindUserByTel(string) (*model.User, error)
	// FindUserByIdNumber 根据身份证号码称查找用户信息
	FindUserByIdNumber(string) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(*model.User) (uint, error)
	// DeleteUserByID 根据用户ID删除用户
	DeleteUserByID(uint) error
	// UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	// FindAll 查找所有用
	FindAll() ([]model.User, error)
}

// NewUserRepository 创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

// FindUserByID 根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID uint) (user *model.User, err error) {
	user = &model.User{}
	u.mysqlDb.First(user, userID)
	return user, nil
}

// FindUserByTel 根据用户名称查找用户信息
func (u *UserRepository) FindUserByTel(tel string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Table("user").Where("tel = ?", tel).First(user).Error
}

// FindUserByIdNumber 根据身份证号码称查找用户信息
func (u *UserRepository) FindUserByIdNumber(idNo string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("id_number = ?", idNo).Find(user).Error
}

// CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID uint, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

// DeleteUserByID 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID uint) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

// UpdateUser 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

// FindAll 查找所有用戶
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
