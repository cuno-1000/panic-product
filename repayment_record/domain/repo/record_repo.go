package repo

import (
	"bytes"
	"github.com/cuno-1000/panic-product/repayment_record/domain/model"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type IRecordRepository interface {
	// InitTable 初始化数据表
	InitTable() error

	AddRecords(records []*model.Record) error
	// FindRecordByID 根据用户ID查找还款记录信息
	FindRecordByID(uint) (*model.Record, error)
	// FindRecordByUUID 根据用户UUID查找还款记录信息
	FindRecordByUUID(string) (*model.Record, error)
	// FindRecordsByUserId 根据用户ID查找还款记录信息
	FindRecordsByUserId(uint) ([]model.Record, error)

	RecordNotAcceptedToPanicEvent(timeWithin time.Time, exp string, maxTimes uint) (record []model.Record, _ error)

	AcceptedToPanicEventByUserId(timeWithin time.Time, userId uint, exp string, maxTimes uint) (bool, error)
	//// FindRecordNotOnTime 查找还款逾期记录信息
	//FindRecordNotOnTime(uint, ...int) (*model.Record, error)
	//// FindRecordByAmount 根据金额查找还款记录信息
	//FindRecordByAmount(string, decimal.Decimal) (records []model.Record, _ error)
	//// CountRecordNotOnTime 根据金额查找还款记录信息
	//CountRecordNotOnTime(int64)
	//// CreateRepaymentRecord 创建还款记录
	//CreateRepaymentRecord(*model.Record) (uint, error)
	//// DeleteRecordByID 根据还款记录ID删除还款记录
	//DeleteRecordByID(uint) error
	//// UpdateRecord 更新还款记录信息
	//UpdateRecord(*model.Record) error
}

// NewRecordRepository 创建RecordRepository
func NewRecordRepository(db *gorm.DB) IRecordRepository {
	return &RecordRepository{mysqlDb: db}
}

type RecordRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (r *RecordRepository) InitTable() error {
	return r.mysqlDb.CreateTable(&model.Record{}).Error
}

//func (r *RecordRepository) AddRecords(records []*model.Record) error {
//	return r.mysqlDb.Create(records).Error
//}
func recordInsertSQL(record *model.Record) string {
	var t int
	if record.IsOnTime {
		t = 1
	} else {
		t = 0
	}
	return "('" + time.Now().Format("2006-01-02 15:04:05") + "'," +
		"'" + record.DueRepaymentTimestamp.Format("2006-01-02 15:04:05") + "'," +
		"'" + record.ActualRepaymentTimestamp.Format("2006-01-02 15:04:05") + "'," +
		strconv.Itoa(record.DateDifference) + "," +
		"" + strconv.Itoa(t) + "," +
		"'" + record.OwedAmount.String() + "'," +
		"" + strconv.FormatUint(uint64(record.UserId), 10) + ")"
}

func (r *RecordRepository) AddRecords(records []*model.Record) error {
	var buffer bytes.Buffer
	sql := "INSERT INTO `record` (`created_at`,`due_repayment_timestamp`,`actual_repayment_timestamp`,`date_difference`,`is_on_time`,`owed_amount`,`user_id`) VALUES"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	a := len(records)
	for i, e := range records {
		if i == a-1 {
			buffer.WriteString(recordInsertSQL(e) + ";")
		} else {
			buffer.WriteString(recordInsertSQL(e) + ",")
		}
	}
	return r.mysqlDb.Exec(buffer.String()).Error
}

// FindRecordByID 根据用户ID查找还款记录信息
func (r *RecordRepository) FindRecordByID(recordID uint) (record *model.Record, _ error) {
	record = &model.Record{}
	return record, r.mysqlDb.First(record, recordID).Error
}

// FindRecordByUUID 根据用户UUID查找还款记录信息
func (r *RecordRepository) FindRecordByUUID(uuid string) (record *model.Record, _ error) {
	record = &model.Record{}
	return record, r.mysqlDb.Where("uuid = ?", uuid).First(record).Error
}

// FindRecordsByUserId 根据用户ID查找还款记录信息
func (r *RecordRepository) FindRecordsByUserId(userId uint) (records []model.Record, _ error) {
	return records, r.mysqlDb.Where("user_id = ?", userId).Find(&records).Error
}

func (r *RecordRepository) RecordNotAcceptedToPanicEvent(timeWithin time.Time, exp string, maxTimes uint) (record []model.Record, _ error) {
	var st string
	t := timeWithin.Format("2006-01-02 15:04:05")
	if exp == "" {
		st = "SELECT user_id FROM record WHERE is_on_time = FALSE AND created_at > \" " + t + " \" GROUP BY user_id HAVING COUNT(is_on_time) > ?;"
		r.mysqlDb.Raw(st, maxTimes).Scan(&record)
	} else {
		st = "SELECT user_id FROM record WHERE is_on_time = FALSE AND created_at > \" " + t + " \" AND " +
			"NOT(" + exp + ") GROUP BY user_id HAVING COUNT(is_on_time) > ?;"
		r.mysqlDb.Raw(st, maxTimes).Scan(&record)
		//r.mysqlDb.Raw(st, exp, maxTimes).Scan(&record)
	}
	return record, nil
}

func (r *RecordRepository) AcceptedToPanicEventByUserId(timeWithin time.Time, userId uint, exp string, maxTimes uint) (bool, error) {
	var record model.Record

	t := timeWithin.Format("2006-01-02 15:04:05")
	var st string
	if exp == "" {
		st = "SELECT user_id FROM record WHERE is_on_time = FALSE AND user_id = ? AND created_at > \" " + t + " \" GROUP BY user_id HAVING COUNT(is_on_time) > ? Limit 1;"
		r.mysqlDb.Raw(st, userId, maxTimes).Scan(&record)
	} else {
		st = "SELECT user_id FROM record WHERE is_on_time = FALSE AND user_id = ? AND created_at > \" " + t + " \" AND NOT(" + exp + ") GROUP BY user_id HAVING COUNT(is_on_time) > ? Limit 1;"
		r.mysqlDb.Raw(st, userId, maxTimes).Scan(&record)
	}
	if record.UserId == userId {
		return true, nil
	}
	return false, nil
}

//func (r *RecordRepository) FetchAllUserIdInRepaymentRecord(before int, times uint8, ownedAmount decimal.Decimal) {
//	now := time.Now()
//	afterTime := now.AddDate(-1*before, 0, 0)
//	a := afterTime.Format("2006-01-02 15:04:05")
//	b := r.mysqlDb.Model(&model.Record{}).Where("create_at > ?", a).Group("user_id").Having("COUNT(*) > ? Where is_on_time <> TRUE")
//	return
//}
