package dto

import (
	"time"

	"github.com/samuelbeaulieu1/gimlet"
)

type ClockDTO struct {
	ID         string
	EmployeeID string    `json:"employee_id"`
	Date       time.Time `json:"date"`
}

func (dto *ClockDTO) GetNewInstance() gimlet.DTO {
	return &ClockDTO{}
}
