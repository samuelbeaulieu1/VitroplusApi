package dto

import "github.com/samuelbeaulieu1/gimlet"

type BranchDTO struct {
	ID      string `json:"id"`
	Store   string `json:"store"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (dto *BranchDTO) GetNewInstance() gimlet.DTO {
	return &BranchDTO{}
}
