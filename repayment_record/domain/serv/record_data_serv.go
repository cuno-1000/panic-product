package serv

import (
	"github.com/cuno-1000/panic-product/repayment_record/domain/model"
	"github.com/cuno-1000/panic-product/repayment_record/domain/repo"
	"github.com/shopspring/decimal"
	"math/rand"
	"time"
)

type IRecordDataService interface {
	AddRecord(userId uint) error
	GetUserIDNotAcceptedToPanicEvent(timeUpperLimit time.Time, exp string, maxTimes uint) ([]uint, error)
	IsAcceptedToPanicEventByUserId(timeUpperLimit time.Time, userId uint64, exp string, maxTimes uint) (bool, error)
}

func NewRecordDataService(recordRepository repo.IRecordRepository) IRecordDataService {
	return &RecordDataService{RecordRepository: recordRepository}
}

type RecordDataService struct {
	RecordRepository repo.IRecordRepository
}

func (r *RecordDataService) AddRecord(userId uint) error {
	var records []*model.Record
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		randomTime1 := rand.Int63n(time.Now().Unix()-94608000) + 94608000
		randomNow1 := time.Unix(randomTime1, 0)
		randomNum := rand.Intn(12+4) - 4
		d, _ := time.ParseDuration("-24h")
		randomNow2 := randomNow1.Add(time.Duration(randomNum) * d)
		Amount := rand.Intn(50000)
		var record *model.Record
		record = &model.Record{
			UserId:                   userId,
			DueRepaymentTimestamp:    randomNow1,
			ActualRepaymentTimestamp: randomNow2,
			DateDifference:           -1 * randomNum,
			IsOnTime:                 !randomNow1.Before(randomNow2),
			OwedAmount:               decimal.NewFromInt(int64(Amount)),
		}
		records = append(records, record)
	}
	err := r.RecordRepository.AddRecords(records)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecordDataService) GetUserIDNotAcceptedToPanicEvent(timeUpperLimit time.Time, exp string, maxTimes uint) ([]uint, error) {
	var userId []uint
	//now := time.Now()
	//setTime := now.AddDate(0, -1*int(monthNumberWithin), 0)
	user, err := r.RecordRepository.RecordNotAcceptedToPanicEvent(timeUpperLimit, exp, maxTimes)
	if err != nil {
		return nil, err
	}
	for _, v := range user {
		userId = append(userId, v.UserId)
	}
	return userId, nil
}

func (r *RecordDataService) IsAcceptedToPanicEventByUserId(timeUpperLimit time.Time, userId uint64, exp string, maxTimes uint) (bool, error) {
	return r.RecordRepository.AcceptedToPanicEventByUserId(timeUpperLimit, uint(userId), exp, maxTimes)
}
