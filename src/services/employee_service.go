package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/gimlet/validators"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
	"github.com/samuelbeaulieu1/vitroplus-api/src/validation"
)

type EmployeeService struct {
	*gimlet.Service[models.EmployeeModel]
}

func NewEmployeeService() *EmployeeService {
	employeeService := &EmployeeService{
		gimlet.NewService[models.EmployeeModel](),
	}
	employeeService.ServiceHandler = employeeService

	return employeeService
}

func (service *EmployeeService) GetEntity() gimlet.Entity[models.EmployeeModel] {
	return entities.NewEmployee()
}

func (service *EmployeeService) RegisterValidators(validator *validators.Validator) {
	validator.RegisterValidation("isValidPin", validation.IsValidPin)
	validator.RegisterValidation("isUniquePin", validation.IsUniquePin)
	validator.RegisterValidation("isValidEmail", validation.IsValidEmail)
	validator.RegisterValidation("isValidBranch", validation.IsValidBranch)
}

func (service *EmployeeService) GetEmployeeDailyReport() responses.Error {
	return nil
}

func (service *EmployeeService) GetEmployeeFromPin(pin string) (*models.EmployeeModel, responses.Error) {
	employee, err := entities.NewEmployee().GetEmployeeFromPin(pin)
	if err != nil {
		return nil, responses.NewError("Aucun employé n'existe avec le pin")
	}

	return employee, nil
}

func (service *EmployeeService) BeforeCreate(request *models.EmployeeModel) responses.Error {
	if request.Pin == "" {
		request.Pin = service.generatePin()
	}

	return nil
}

func (service *EmployeeService) generatePin() string {
	employee := entities.NewEmployee()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var pin string
	for {
		pin = ""
		for i := 0; i < models.PinLength; i++ {
			pin += fmt.Sprintf("%d", r.Intn(9))
		}
		if !employee.PinExists(pin) {
			break
		}
	}

	return pin
}
