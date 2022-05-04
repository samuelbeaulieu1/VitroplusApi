package dto

import (
	"time"

	"github.com/samuelbeaulieu1/gimlet"
)

type EmployeeClocksDTO struct {
	Clocks         []ClockDTO                 `json:"clocks"`
	Date           time.Time                  `json:"date"`
	TotalTime      string                     `json:"total_time"`
	TotalTimeValue int                        `json:"total_time_value"`
	TotalTimeFloat float64                    `json:"total_time_float"`
	Employee       *EmployeeIdentificationDTO `json:"employee"`
}

func (dto *EmployeeClocksDTO) GetNewInstance() gimlet.DTO {
	return &EmployeeClocksDTO{}
}
