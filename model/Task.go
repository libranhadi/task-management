package model

import (
	"errors"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func (tr *TaskRequest) ToTask() (Task, error) {
	startDate, err := time.Parse("2006-01-02", tr.StartDate)
	if err != nil {
		return Task{}, errors.New("invalid start date format")
	}
	endDate, err := time.Parse("2006-01-02", tr.EndDate)
	if err != nil {
		return Task{}, errors.New("invalid end date format")
	}
	if endDate.Before(startDate) {
		return Task{}, errors.New("end date must be after start date")
	}
	return Task{
		Title:       tr.Title,
		Description: tr.Description,
		Status:      tr.Status,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
