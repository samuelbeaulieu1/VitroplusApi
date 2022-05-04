package classes

import "time"

type ClocksReport map[int64]*ClocksDayReport

type ClocksDayReport struct {
	Date   time.Time                  `json:"date"`
	Report map[string]*EmployeeClocks `json:"report"`
}
