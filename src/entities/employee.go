package entities

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dao"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type Employee struct {
	dao *dao.EmployeeDao
}

func NewEmployee() *Employee {
	return &Employee{
		dao.NewEmployeeDao(),
	}
}

func (employee *Employee) Delete(id string) error {
	return employee.dao.Delete(id, &models.EmployeeModel{})
}

func (employee *Employee) Update(id string, request *models.EmployeeModel) error {
	record := &models.EmployeeModel{ID: id}
	return employee.dao.Update(record, request)
}

func (employee *Employee) Create(request *models.EmployeeModel) (*models.EmployeeModel, error) {
	employeeData := &models.EmployeeModel{
		ID:              gimlet.CreateNewID(employee.dao, &models.EmployeeModel{}, gimlet.DefaultIDLength),
		Firstname:       request.Firstname,
		Lastname:        request.Lastname,
		Email:           request.Email,
		IsConstantHours: request.IsConstantHours,
		ConstantHours:   request.ConstantHours,
		BranchID:        request.BranchID,
		Pin:             request.Pin,
	}
	err := employee.dao.Create(employeeData)
	return employeeData, err
}

func (employee *Employee) Get(id string) (*models.EmployeeModel, error) {
	record := &models.EmployeeModel{}
	err := employee.dao.Get(id, record)
	return record, err
}

func (employee *Employee) GetAll() (*[]models.EmployeeModel, error) {
	return employee.dao.GetAll()
}

func (employee *Employee) GetEmployeeFromPin(pin string) (*models.EmployeeModel, error) {
	return employee.dao.GetEmployeeFromPin(pin)
}

func (employee *Employee) PinExists(pin string) bool {
	return employee.dao.PinExists(pin)
}

func (branch *Employee) Exists(id string) bool {
	return branch.dao.ExistsByID(id, &models.EmployeeModel{})
}
