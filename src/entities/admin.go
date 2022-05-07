package entities

import (
	"github.com/samuelbeaulieu1/vitroplus-api/src/dao"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	dao *dao.AdminDao
}

func NewAdmin() *Admin {
	return &Admin{
		dao.NewAdminDao(),
	}
}

func (admin *Admin) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (admin *Admin) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (admin *Admin) Update(newPassword string) error {
	hash, err := admin.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return admin.dao.Update(&models.AdminModel{
		ID: 1,
	}, &models.AdminModel{
		Password: hash,
	})
}

func (admin *Admin) Get() (*models.AdminModel, error) {
	record := &models.AdminModel{}
	err := admin.dao.Get("1", record)
	return record, err
}
