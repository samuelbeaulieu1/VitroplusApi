package dto

import (
	"time"
)

type BranchReportDTO struct {
	StartDate       time.Time                           `json:"start_date"`
	EndDate         time.Time                           `json:"end_date"`
	Branch          *BranchDTO                          `json:"branch"`
	EmployeesReport map[string]*BranchEmployeeReportDTO `json:"employees_report"`
}

type BranchEmployeeReportDTO struct {
	Employee  *EmployeeDTO `json:"employee"`
	TotalTime float64      `json:"total_time"`
	Overtime  float64      `json:"overtime"`
}
