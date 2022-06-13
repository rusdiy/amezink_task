package model

import "time"

type Record struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Marks     []Mark    `json:"marks"`
	CreatedAt time.Time `json:"created_at"`
}

type RecordResponses struct {
	ID         int64  `json:"id"`
	CreatedAt  string `json:"created_at"`
	TotalMarks int    `json:"totalMarks"`
}

type RecordRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	MinCount  int    `json:"min_count"`
	MaxCount  int    `json:"max_count"`
}

type Mark struct {
	ID        int64     `json:"id"`
	Mark      int       `json:"mark"`
	RecordID  int64     `json:"record_id"`
	CreatedAt time.Time `json:"created_at"`
}
