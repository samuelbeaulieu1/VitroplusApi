package classes

import (
	"time"

	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type UpdateEmployeeClocksRequest struct {
	Clocks     []models.ClockModel `json:"clocks"`
	EmployeeID string
	Date       time.Time
}
