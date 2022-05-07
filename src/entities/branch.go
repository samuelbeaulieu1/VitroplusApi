package entities

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dao"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

type Branch struct {
	dao *dao.BranchDao
}

func NewBranch() *Branch {
	return &Branch{
		dao.NewBranchDao(),
	}
}

func (branch *Branch) Delete(id string) error {
	return branch.dao.Delete(id, &models.BranchModel{})
}

func (branch *Branch) Update(id string, request *models.BranchModel) error {
	record := &models.BranchModel{ID: id}
	return branch.dao.Update(record, request)
}

func (branch *Branch) Create(request *models.BranchModel) (*models.BranchModel, error) {
	branchData := &models.BranchModel{
		ID:      gimlet.CreateNewID(branch.dao, &models.BranchModel{}, gimlet.DefaultIDLength),
		Store:   request.Store,
		Address: request.Address,
		Phone:   request.Phone,
		Owner:   request.Owner,
	}
	err := branch.dao.Create(branchData)
	return branchData, err
}

func (branch *Branch) Get(id string) (*models.BranchModel, error) {
	record := &models.BranchModel{}
	err := branch.dao.Get(id, record)
	return record, err
}

func (branch *Branch) GetAll() (*[]models.BranchModel, error) {
	return branch.dao.GetAll()
}

func (branch *Branch) GetEmployees(branchID string) ([]*models.EmployeeModel, error) {
	return branch.dao.GetEmployees(branchID)
}

func (branch *Branch) Exists(id string) bool {
	return branch.dao.ExistsByID(id, &models.BranchModel{})
}
