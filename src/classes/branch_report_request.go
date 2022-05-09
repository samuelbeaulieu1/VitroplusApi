package classes

import (
	"time"

	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type BranchReportRequest struct {
	StartDate       time.Time
	EndDate         time.Time
	Branch          *models.BranchModel
	EmployeesReport []*BranchEmployeeReport `json:"employees_report"`
	Logo            string
}
