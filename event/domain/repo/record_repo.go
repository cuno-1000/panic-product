package repo

import "github.com/jinzhu/gorm"

type IRecordRepository interface {
}

// NewRecordRepository 创建RecordRepository
func NewRecordRepository(db *gorm.DB) IRecordRepository {
	return &RecordRepository{mysqlDb: db}
}

type RecordRepository struct {
	mysqlDb *gorm.DB
}
