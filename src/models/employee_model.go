package models

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"gorm.io/gorm"
)

const (
	PinLength = 4
)

type EmployeeModel struct {
	gorm.Model
	ID              string
	Firstname       string  `json:"first_name" validate:"required" label:"Prénom"`
	Lastname        string  `json:"last_name" validate:"required" label:"Nom"`
	Pin             string  `json:"pin" validate:"requiredOnUpdate,isValidPin,isUniquePin" label:"Pin"`
	BranchID        string  `json:"branch_id" validate:"required,isValidBranch" label:"Succursale"`
	IsConstantHours bool    `json:"is_constant_hours"`
	ConstantHours   float64 `json:"constant_hours"`
	Email           string  `json:"email" validate:"isValidEmail" label:"Courriel"`
}

func (model EmployeeModel) TableName() string {
	return "vp_employee"
}

func (model EmployeeModel) ToDTO() gimlet.DTO {
	var modelDto *dto.EmployeeDTO
	gimlet.ParseModelToDTO(model, &modelDto)

	return modelDto
}
