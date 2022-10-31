package serv

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/model"
	"github.com/cuno-1000/panic-product/repayment_record/domain/repo"
	"math/rand"
	"time"
)

type IBlacklistDataService interface {
	AppendBlacklist(uint) error
	FetchUserIdInBlacklist() ([]uint, error)
	IsUserIdInBlacklist(userId uint64) (bool, error)
}

func NewBlacklistDataService(blacklistRepository repo.IBlacklistRepository) IBlacklistDataService {
	return &BlacklistDataService{BlacklistRepository: blacklistRepository}
}

type BlacklistDataService struct {
	BlacklistRepository repo.IBlacklistRepository
}

var reason []string

func init() {
	reason = append(reason,
		"畏罪潜逃",
		"欠下巨款",
		"偷税漏税",
	)
}

func (b *BlacklistDataService) AppendBlacklist(userId uint) error {
	var black *model.BlackList
	rand.Seed(time.Now().Unix())
	black = &model.BlackList{UserId: userId, Reason: reason[rand.Intn(3)]}
	if err := b.BlacklistRepository.AppendBlacklist(black); err != nil {
		return err
	}
	return nil
}

func (b *BlacklistDataService) FetchUserIdInBlacklist() ([]uint, error) {
	return b.BlacklistRepository.FetchAllUserId()
}

func (b *BlacklistDataService) IsUserIdInBlacklist(userId uint64) (bool, error) {
	return b.BlacklistRepository.IsUserIdExist(uint(userId))
}
