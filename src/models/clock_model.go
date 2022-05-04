package models

import (
	"time"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"gorm.io/gorm"
)

type ClockModel struct {
	gorm.Model
	ID         string
	EmployeeID string    `json:"employee_id"`
	Date       time.Time `json:"date"`
}

func (model ClockModel) TableName() string {
	return "vp_clock"
}

func (model ClockModel) ToDTO() gimlet.DTO {
	var modelDto *dto.ClockDTO
	gimlet.ParseModelToDTO(model, &modelDto)

	return modelDto
}
