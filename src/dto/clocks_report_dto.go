package dto

import "time"

type ClocksReportDTO map[int64]*ClocksDayReportDTO

type ClocksDayReportDTO struct {
	Date   time.Time                     `json:"date"`
	Report map[string]*EmployeeClocksDTO `json:"report"`
}
