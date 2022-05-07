package services

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/actions"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/gimlet/validators"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
	"github.com/samuelbeaulieu1/vitroplus-api/src/validation"
)

type BranchService struct {
	*gimlet.Service[models.BranchModel]
}

func NewBranchService() *BranchService {
	branchService := &BranchService{
		gimlet.NewService[models.BranchModel](),
	}
	branchService.ServiceHandler = branchService

	return branchService
}

func (service *BranchService) GetEntity() gimlet.Entity[models.BranchModel] {
	return entities.NewBranch()
}

func (service *BranchService) RegisterValidators(action actions.Action, request *models.BranchModel, validator *validators.Validator) {
	validator.RegisterValidation("isValidPhone", validation.IsValidPhone)
}

func (service *BranchService) GetBranchEmployees(branchID string) ([]*models.EmployeeModel, responses.Error) {
	if err := service.Exists(branchID); err != nil {
		return nil, responses.NewError("Succursale inexistante")
	}

	employees, err := entities.NewBranch().GetEmployees(branchID)
	if err != nil {
		return nil, responses.NewError("Une erreur inattendue est survenue en récupérant les employés")
	}

	return employees, nil
}
