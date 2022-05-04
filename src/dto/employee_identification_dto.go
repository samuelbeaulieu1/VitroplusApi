package dto

import "github.com/samuelbeaulieu1/gimlet"

type EmployeeIdentificationDTO struct {
	ID        string  `json:"id"`
	Firstname string  `json:"first_name"`
	Lastname  string  `json:"last_name"`
	BranchID  string  `json:"branch_id"`
	Email     *string `json:"email"`
}

func (dto *EmployeeIdentificationDTO) GetNewInstance() gimlet.DTO {
	return &EmployeeIdentificationDTO{}
}
