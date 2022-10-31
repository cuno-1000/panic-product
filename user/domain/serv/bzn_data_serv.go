package serv

import (
	"github.com/cuno-1000/panic-product/user/domain/model"
	"github.com/cuno-1000/panic-product/user/domain/repo"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

type IBznDataService interface {
	// AgeCounter 根据身份证号码得知年龄
	AgeCounter(userIdNo string) (age uint16)
	// GetGender 根据身份证号码得知性别
	GetGender(userIdNo string) (gender string)
	// ValidateIdCartNumber 验证身份证号码
	ValidateIdCartNumber(userIdNo string) bool

	FetchNormalAdult() ([]model.User, error)

	ReduceUserBalance(userId uint, amount string) (bool, error)
}

// NewBznDataService 创建实例
func NewBznDataService(userRepository repo.IUserRepository) IBznDataService {
	return &BznDataService{BznRepository: userRepository}
}

type BznDataService struct {
	BznRepository repo.IUserRepository
}

// AgeCounter 根据身份证号码得知年龄
func (b *BznDataService) AgeCounter(userIdNo string) (age uint16) {
	birthYear, _ := strconv.Atoi(userIdNo[6:10])
	birthMouth, _ := strconv.Atoi(userIdNo[10:12])
	birthDay, _ := strconv.Atoi(userIdNo[12:14])
	now := time.Now()
	age = uint16(now.Year() - birthYear)
	if int(now.Month()) < birthMouth ||
		int(now.Month()) == birthMouth && now.Day() < birthDay {
		age--
	}
	return
}

// GetGender 根据身份证号码得知性别
func (b *BznDataService) GetGender(userIdNo string) (gender string) {
	var genderSelect = map[bool]string{
		true:  "男",
		false: "女",
	}
	if t, _ := strconv.Atoi(string(userIdNo[16])); t%2 == 1 {
		return genderSelect[true]
	} else {
		return genderSelect[false]
	}
}

// ValidateIdCartNumber 验证身份证号码
func (b *BznDataService) ValidateIdCartNumber(userIdNo string) bool {
	if len(userIdNo) != 18 {
		return false
	}
	return GetIdNumberValidateCODE(userIdNo[0:17]) == userIdNo[17:18]
}

// getIdNumberValidateCODE
func GetIdNumberValidateCODE(id17 string) string {
	if len(id17) != 17 {
		return ""
	}
	var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var validate = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	var mode, sum, t int
	for i := 0; i < len(id17); i++ {
		t, _ = strconv.Atoi(string(id17[i]))
		sum += t * weight[i]
	}
	mode = sum % 11
	return string(validate[mode])
}

func (b *BznDataService) FetchNormalAdult() ([]model.User, error) {
	return b.BznRepository.FindAll()
}

func (b *BznDataService) ReduceUserBalance(userId uint, amount string) (bool, error) {
	a, err := decimal.NewFromString(amount)
	if err != nil {
		return false, err
	}
	user, err := b.BznRepository.FindUserByID(userId)
	if user.AccountBalance.Cmp(a) == -1 {
		return false, nil
	}

	user.AccountBalance = user.AccountBalance.Sub(a)
	return true, b.BznRepository.UpdateUser(user)
}
