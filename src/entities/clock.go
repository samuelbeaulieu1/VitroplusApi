package entities

import (
	"time"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dao"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type Clock struct {
	dao *dao.ClockDao
}

func NewClock() *Clock {
	return &Clock{
		dao.NewClockDao(),
	}
}

func (clock *Clock) Delete(id string) error {
	return clock.dao.Delete(id, &models.ClockModel{})
}

func (clock *Clock) Exists(id string) bool {
	return clock.dao.ExistsByID(id, &models.ClockModel{})
}

func (clock *Clock) UpdateEmployeeClocks(req *classes.UpdateEmployeeClocksRequest, clocks *[]models.ClockModel) error {
	newClocks := []models.ClockModel{}
	year, month, day := req.Date.Date()
	for _, c := range *clocks {
		clockDate := c.Date.In(time.Local)
		date := time.Date(year, month, day, clockDate.Hour(), clockDate.Minute(), 0, 0, time.Local)
		newClocks = append(newClocks, models.ClockModel{
			ID:         gimlet.CreateNewID(clock.dao, &models.ClockModel{}, gimlet.DefaultIDLength),
			Date:       date,
			EmployeeID: req.EmployeeID,
		})
	}

	return clock.dao.CreateMultipleClocks(&newClocks)
}

func (clock *Clock) GetEmployeeClocksBetween(employeeID string, startDate time.Time, endDate time.Time) (*[]models.ClockModel, error) {
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.Local)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, -1, time.Local)

	return clock.dao.GetEmployeeClocks(employeeID, startDate, endDate)
}

func (clock *Clock) GetBranchClocksBetween(branchID string, startDate time.Time, endDate time.Time) (*[]models.ClockModel, error) {
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.Local)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, -1, time.Local)

	return clock.dao.GetBranchClocks(branchID, startDate, endDate)
}

func (clock *Clock) GetEmployeeClocks(employeeID string, date time.Time) (*[]models.ClockModel, error) {
	year, month, day := date.Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	endDate := time.Date(year, month, day+1, 0, 0, 0, -1, time.Local)

	return clock.dao.GetEmployeeClocks(employeeID, startDate, endDate)
}

func (clock *Clock) GetLastEmployeeClock(employeeID string) (*models.ClockModel, error) {
	return clock.dao.GetLastEmployeeClock(employeeID)
}

func (clock *Clock) ClockIn(employeeID string, date time.Time) error {
	return clock.dao.Create(&models.ClockModel{
		ID:         gimlet.CreateNewID(clock.dao, &models.ClockModel{}, gimlet.DefaultIDLength),
		Date:       date,
		EmployeeID: employeeID,
	})
}
