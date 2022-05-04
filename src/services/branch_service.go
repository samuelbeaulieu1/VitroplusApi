package services

import (
	"github.com/samuelbeaulieu1/gimlet"
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

func (service *BranchService) RegisterValidators(validator *validators.Validator) {
	validator.RegisterValidation("isValidPhone", validation.IsValidPhone)
}
