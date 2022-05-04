package dto

import (
	"time"
)

type ClockEntryDTO struct {
	Employee string    `json:"employee"`
	Date     time.Time `json:"date"`
}
