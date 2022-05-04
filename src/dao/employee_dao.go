package dao

import (
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type EmployeeDao struct {
	Dao
}

func NewEmployeeDao() *EmployeeDao {
	dao := &EmployeeDao{}
	connectDb(&dao.Dao)

	return dao
}

func (dao *EmployeeDao) GetAll() (*[]models.EmployeeModel, error) {
	var employees []models.EmployeeModel
	res := dao.Db.Model(&models.EmployeeModel{}).Order("firstname, lastname ASC").Find(&employees)

	return &employees, res.Error
}

func (dao *EmployeeDao) GetEmployeeFromPin(pin string) (*models.EmployeeModel, error) {
	var employee models.EmployeeModel
	res := dao.Db.Model(&models.EmployeeModel{}).Where("pin = ?", pin).First(&employee)

	return &employee, res.Error
}

func (dao *EmployeeDao) PinExists(pin string) bool {
	var exists bool
	res := dao.Db.Model(&models.EmployeeModel{}).Select("count(id) > 0").Where("pin = ?", pin).Find(&exists)

	return res.Error == nil && exists
}
