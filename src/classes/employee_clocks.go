package classes

import (
	"time"

	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type EmployeeClocks struct {
	Clocks         []models.ClockModel
	Date           time.Time             `json:"date"`
	TotalTime      string                `json:"total_time"`
	TotalTimeValue int                   `json:"total_time_value"`
	TotalTimeFloat float64               `json:"total_time_float"`
	Employee       *models.EmployeeModel `json:"employee"`
}
