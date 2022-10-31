package repo

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/model"
	"github.com/jinzhu/gorm"
)

type IBlacklistRepository interface {
	// InitTable 初始化数据表
	InitTable() error

	AppendBlacklist(*model.BlackList) error
	// FetchAllUserId 用于预热
	FetchAllUserId() ([]uint, error)
	IsUserIdExist(userId uint) (bool, error)
}

// NewBlacklistRepository 创建BlacklistRepository
func NewBlacklistRepository(db *gorm.DB) IBlacklistRepository {
	return &BlacklistRepository{mysqlDb: db}
}

type BlacklistRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (b *BlacklistRepository) InitTable() error {
	return b.mysqlDb.CreateTable(&model.BlackList{}).Error
}

func (b *BlacklistRepository) AppendBlacklist(black *model.BlackList) error {
	return b.mysqlDb.Create(black).Error
}

func (b *BlacklistRepository) FetchAllUserId() (blacklistItems []uint, err error) {
	var blacklist []model.BlackList
	err = b.mysqlDb.Select("user_id").Limit(10).Find(&blacklist).Error
	for _, v := range blacklist {
		blacklistItems = append(blacklistItems, v.UserId)
	}
	//b.mysqlDb.Raw("SELECT user_id FROM blacklist").Scan(&blacklistItems)
	return blacklistItems, err
}

func (b *BlacklistRepository) IsUserIdExist(userId uint) (bool, error) {
	var blacklist = &model.BlackList{}
	st := "SELECT user_id FROM black_list WHERE user_id = ? Limit 1;"
	b.mysqlDb.Raw(st, userId).Scan(blacklist)
	if blacklist.UserId == userId {
		return true, nil
	}
	return false, nil
}
