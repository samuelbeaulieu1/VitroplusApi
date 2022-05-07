package classes

import (
	"time"

	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type BranchReport struct {
	StartDate       time.Time                        `json:"start_date"`
	EndDate         time.Time                        `json:"end_date"`
	Branch          *models.BranchModel              `json:"branch"`
	EmployeesReport map[string]*BranchEmployeeReport `json:"employees_report"`
	Logo            string
}

type BranchEmployeeReport struct {
	Employee  *models.EmployeeModel `json:"employee"`
	TotalTime float64               `json:"total_time"`
	Overtime  float64               `json:"overtime"`
}
