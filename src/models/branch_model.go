package models

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"gorm.io/gorm"
)

type BranchModel struct {
	gorm.Model
	ID      string
	Store   string `json:"store" validate:"required" label:"Magasin"`
	Owner   string `json:"owner" validate:"required" label:"Propriétaire"`
	Address string `json:"address" validate:"required" label:"Adresse"`
	Phone   string `json:"phone" validate:"required,isValidPhone" label:"Téléphone"`
}

func (model BranchModel) TableName() string {
	return "vp_branch"
}

func (model BranchModel) ToDTO() gimlet.DTO {
	var modelDto *dto.BranchDTO
	gimlet.ParseModelToDTO(model, &modelDto)

	return modelDto
}
