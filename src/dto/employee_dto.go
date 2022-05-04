package dto

import "github.com/samuelbeaulieu1/gimlet"

type EmployeeDTO struct {
	ID              string  `json:"id"`
	Firstname       string  `json:"first_name"`
	Lastname        string  `json:"last_name"`
	Pin             string  `json:"pin"`
	BranchID        string  `json:"branch_id"`
	IsConstantHours bool    `json:"is_constant_hours"`
	ConstantHours   float64 `json:"constant_hours"`
	Email           *string `json:"email"`
}

func (dto *EmployeeDTO) GetNewInstance() gimlet.DTO {
	return &EmployeeDTO{}
}
