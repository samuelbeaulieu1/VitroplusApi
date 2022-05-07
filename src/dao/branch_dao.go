package dao

import (
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type BranchDao struct {
	Dao
}

func NewBranchDao() *BranchDao {
	dao := &BranchDao{}
	connectDb(&dao.Dao)

	return dao
}

func (dao *BranchDao) GetAll() (*[]models.BranchModel, error) {
	var branches []models.BranchModel

	res := dao.Db.Model(&models.BranchModel{}).Order("store ASC").Find(&branches)

	return &branches, res.Error
}

func (dao *BranchDao) GetEmployees(branchID string) ([]*models.EmployeeModel, error) {
	var employees []*models.EmployeeModel

	res := dao.Db.Model(&models.EmployeeModel{}).Where("branch_id = ?", branchID).Find(&employees)

	return employees, res.Error
}
