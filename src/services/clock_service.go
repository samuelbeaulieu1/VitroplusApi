package services

import (
	"fmt"
	"math"
	"time"

	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

const (
	workTimeQuarter      = 15
	workTimeThreshold    = 7.5
	overtimeHrsThreshold = 40
	startOfDay           = 12
)

type ClockService struct{}

func NewClockService() *ClockService {
	clockService := &ClockService{}

	return clockService
}

func timeDurationToString(minutes int) string {
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)
}

func (service *ClockService) GetEmployeeClocks(employeeID string, date time.Time) (*classes.EmployeeClocks, responses.Error) {
	employee, err := NewEmployeeService().Get(employeeID)
	if err != nil {
		return nil, err
	}

	clocks, err := entities.NewClock().GetEmployeeClocks(employeeID, date)
	if err != nil {
		return nil, responses.NewError("Impossible de récupérer les entrées de l'employé")
	}

	totalMinutes := service.calculateClocksTime(clocks)
	return &classes.EmployeeClocks{
		Clocks:         *clocks,
		Date:           date,
		TotalTime:      timeDurationToString(totalMinutes),
		TotalTimeValue: totalMinutes,
		TotalTimeFloat: float64(totalMinutes) / 60.0,
		Employee:       employee,
	}, nil
}

func (service *ClockService) calculateClocksTime(clocks *[]models.ClockModel) int {
	totalMinutes := 0
	firstIndex := 0

	for firstIndex < len(*clocks) {
		clock := (*clocks)[firstIndex]
		if firstIndex+1 < len(*clocks) {
			next := (*clocks)[firstIndex+1]

			diff := next.Date.Sub(clock.Date)
			minutes := int(math.Floor(diff.Minutes()))
			remainder := minutes % workTimeQuarter
			minutes -= remainder
			if float64(remainder) >= workTimeThreshold {
				totalMinutes += minutes + workTimeQuarter
			} else {
				totalMinutes += minutes
			}
		}

		firstIndex += 2
	}

	return totalMinutes
}

func (service *ClockService) GetBranchClocksInTimeframe(branchID string, startDate time.Time, endDate time.Time) (*classes.ClocksReport, responses.Error) {
	if err := NewBranchService().Exists(branchID); err != nil {
		return nil, err
	}
	clocks, err := entities.NewClock().GetBranchClocksBetween(branchID, startDate, endDate)
	if err != nil {
		return nil, responses.NewError("Impossible de récupérer les entrées des employés de la succursale")
	}
	clocksReport := service.groupClocksByDate(clocks)

	return clocksReport, nil
}

func (service *ClockService) GetEmployeeClocksInTimeframe(employeeID string, startDate time.Time, endDate time.Time) (*classes.ClocksReport, responses.Error) {
	if err := NewEmployeeService().Exists(employeeID); err != nil {
		return nil, err
	}
	clocks, err := entities.NewClock().GetEmployeeClocksBetween(employeeID, startDate, endDate)
	if err != nil {
		return nil, responses.NewError("Impossible de récupérer les entrées de l'employé")
	}
	clocksReport := service.groupClocksByDate(clocks)

	return clocksReport, nil
}

func (service *ClockService) groupClocksByDate(clocks *[]models.ClockModel) *classes.ClocksReport {
	clocksReport := make(classes.ClocksReport)
	for _, clock := range *clocks {
		year, month, day := clock.Date.Date()
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		unixTimestamp := date.Unix()
		if _, ok := clocksReport[unixTimestamp]; !ok {
			clocksReport[unixTimestamp] = &classes.ClocksDayReport{
				Date:   date,
				Report: make(map[string]*classes.EmployeeClocks),
			}
		}
		if _, ok := clocksReport[unixTimestamp].Report[clock.EmployeeID]; !ok {
			if employee, err := NewEmployeeService().Get(clock.EmployeeID); err == nil {
				report := clocksReport[unixTimestamp].Report
				report[clock.EmployeeID] = &classes.EmployeeClocks{
					Clocks:   []models.ClockModel{},
					Date:     date,
					Employee: employee,
				}
			}
		}

		entry := clocksReport[unixTimestamp].Report[clock.EmployeeID]
		entry.Clocks = append(entry.Clocks, clock)
	}
	service.calculateReportClocksTime(&clocksReport)

	return &clocksReport
}

func (service *ClockService) calculateReportClocksTime(report *classes.ClocksReport) {
	for _, employeesClocks := range *report {
		for _, employeeClocks := range employeesClocks.Report {
			totalMinutes := service.calculateClocksTime(&employeeClocks.Clocks)
			employeeClocks.TotalTime = timeDurationToString(totalMinutes)
			employeeClocks.TotalTimeValue = totalMinutes
			employeeClocks.TotalTimeFloat = float64(totalMinutes) / 60.0
		}
	}
}

func (service *ClockService) UpdateEmployeeClocks(req *classes.UpdateEmployeeClocksRequest) responses.Error {
	employeeClocks, err := service.GetEmployeeClocks(req.EmployeeID, req.Date)
	if err != nil {
		return err
	}

	clockEntity := entities.NewClock()
	if len(req.Clocks) > 0 {
		err = clockEntity.UpdateEmployeeClocks(req, &req.Clocks)
		if err != nil {
			return responses.NewError("Impossible de sauvegarder les entrées pour l'employé")
		}
	}

	for _, clock := range employeeClocks.Clocks {
		clockEntity.Delete(clock.ID)
	}

	return nil
}

func (service *ClockService) ClockInOut(req *classes.ClockInRequest) (*dto.ClockEntryDTO, responses.Error) {
	employee, err := service.getAndVerifyClockInEmployee(req.Pin)
	if err != nil {
		return nil, err
	}
	date := time.Now().In(time.UTC)
	if date.Hour() < startOfDay {
		date = time.Date(date.Year(), date.Month(), date.Day(), startOfDay, 0, 0, 0, time.UTC)
	}
	if err := service.verifyEmployeeClockIn(employee, date); err != nil {
		return nil, err
	}
	err = entities.NewClock().ClockIn(employee.ID, date)
	if err != nil {
		return nil, responses.NewError("Une erreur est survenue en tentant de sauvegarder l'entrée")
	}

	return &dto.ClockEntryDTO{
		Employee: fmt.Sprintf("%s %s", employee.Firstname, employee.Lastname),
		Date:     date,
	}, nil
}

func (service *ClockService) getAndVerifyClockInEmployee(pin string) (*models.EmployeeModel, responses.Error) {
	employee, err := NewEmployeeService().GetEmployeeFromPin(pin)
	if err != nil {
		return nil, err
	}
	if *employee.IsConstantHours {
		return nil, responses.NewError(fmt.Sprintf("%s %s est à temps constant", employee.Firstname, employee.Lastname))
	}

	return employee, nil
}

func (service *ClockService) verifyEmployeeClockIn(employee *models.EmployeeModel, date time.Time) responses.Error {
	if err := service.verifyEmployeeLastDay(employee, date); err != nil {
		return err
	}
	if err := service.verifyEmployeeLastClockTime(employee, date); err != nil {
		return err
	}

	return nil
}

func (service *ClockService) verifyEmployeeLastDay(employee *models.EmployeeModel, date time.Time) responses.Error {
	if lastClock, err := entities.NewClock().GetLastEmployeeClock(employee.ID); err == nil {
		lastDayClocks, err := service.GetEmployeeClocks(employee.ID, lastClock.Date)
		validClocks := false
		if err == nil {
			year1, month1, day1 := date.Date()
			year2, month2, day2 := lastDayClocks.Date.Date()
			validClocks = (year1 == year2 && month1 == month2 && day1 == day2) || len(lastDayClocks.Clocks)%2 == 0
		}

		if !validClocks {
			message := fmt.Sprintf("%s %s n'a pas complété sa dernière journée de travail. Veuillez contacter un administrateur pour la compléter manuellement", employee.Firstname, employee.Lastname)
			return responses.NewError(message)
		}
	}

	return nil
}

func (service *ClockService) verifyEmployeeLastClockTime(employee *models.EmployeeModel, date time.Time) responses.Error {
	if lastClock, err := entities.NewClock().GetLastEmployeeClock(employee.ID); err == nil {
		diff := date.Sub(lastClock.Date)
		if diff.Minutes() <= workTimeQuarter {
			message := fmt.Sprintf("%s %s a déjà entré son pin il y a moins de %d minutes", employee.Firstname, employee.Lastname, workTimeQuarter)
			return responses.NewError(message)
		}
	}

	return nil
}
