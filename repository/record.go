package repository

import (
	"strconv"

	"github.com/rusdiy/amezink_task/config/database"
	"github.com/rusdiy/amezink_task/model"
)

type IRecordRepository interface {
	GetMarks(string, string, int, int) ([]model.RecordResponses, error)
}

type RecordRepository struct {
	DB database.Database
}

func (r RecordRepository) GetMarks(startDate string, endDate string, minCount int, maxCount int) ([]model.RecordResponses, error) {
	var recordResponses []model.RecordResponses
	query := `
		SELECT
			r.id,
			r.created_at,
			SUM(m.marks) AS totalMarks
		FROM records AS r
		LEFT JOIN marks AS m ON m.record_id = r.id
		WHERE (
				r.created_at >= '` + startDate + `' AND
				r.created_at <= '` + endDate + `'
			)
		GROUP BY r.id
			HAVING (
				SUM(m.marks) >= ` + strconv.Itoa(minCount) + ` AND 
				SUM(m.marks) <= ` + strconv.Itoa(maxCount) + `
			);
	`
	rows, err := r.DB.MySQL.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recordResponse = model.RecordResponses{}
		var err = rows.Scan(
			&recordResponse.ID,
			&recordResponse.CreatedAt,
			&recordResponse.TotalMarks,
		)
		if err != nil {
			return nil, err
		}
		recordResponses = append(recordResponses, recordResponse)
	}
	return recordResponses, nil
}
