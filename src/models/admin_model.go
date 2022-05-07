package models

import (
	"github.com/samuelbeaulieu1/gimlet"
)

type AdminModel struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func (model AdminModel) TableName() string {
	return "vp_administration"
}

func (model AdminModel) ToDTO() gimlet.DTO {
	return nil
}
