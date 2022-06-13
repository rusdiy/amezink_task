package service

import (
	"github.com/rusdiy/amezink_task/model"
	"github.com/rusdiy/amezink_task/repository"
)

type RecordService struct {
	RecordRepo repository.IRecordRepository
}

type IRecordService interface {
	GetMarks(string, string, int, int) ([]model.RecordResponses, error)
}

func (rs RecordService) GetMarks(startDate string, endDate string, minCount int, maxCount int) ([]model.RecordResponses, error) {
	data, err := rs.RecordRepo.GetMarks(startDate, endDate, minCount, maxCount)
	if err != nil {
		return nil, err
	}

	return data, nil
}
