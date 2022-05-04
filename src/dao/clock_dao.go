package dao

import (
	"time"

	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type ClockDao struct {
	Dao
}

func NewClockDao() *ClockDao {
	dao := &ClockDao{}
	connectDb(&dao.Dao)

	return dao
}

func (dao *ClockDao) GetBranchClocks(branchID string, startDate time.Time, endDate time.Time) (*[]models.ClockModel, error) {
	var clocks []models.ClockModel

	employees := dao.Db.Model(&models.EmployeeModel{}).Select("id").Where("branch_id = ?", branchID)
	res := dao.Db.Model(&models.ClockModel{}).
		Where("employee_id IN (?)", employees).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Order("date ASC").Find(&clocks)

	return &clocks, res.Error
}

func (dao *ClockDao) GetEmployeeClocks(employeeID string, startDate time.Time, endDate time.Time) (*[]models.ClockModel, error) {
	var clocks []models.ClockModel
	res := dao.Db.Model(&models.ClockModel{}).
		Where("employee_id = ?", employeeID).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Order("date ASC").Find(&clocks)

	return &clocks, res.Error
}

func (dao *ClockDao) GetLastEmployeeClock(employeeID string) (*models.ClockModel, error) {
	var clock models.ClockModel
	res := dao.Db.Model(&models.ClockModel{}).
		Where("employee_id = ?", employeeID).
		Order("date DESC").
		Limit(1).First(&clock)

	return &clock, res.Error
}

func (dao *ClockDao) CreateMultipleClocks(clocks *[]models.ClockModel) error {
	res := dao.Db.Create(&clocks)

	return res.Error
}
