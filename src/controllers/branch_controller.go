package controllers

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/middlewares"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
	"github.com/samuelbeaulieu1/vitroplus-api/src/services"
)

type BranchController struct {
	*gimlet.Controller[models.BranchModel]
}

func NewBranchController() *BranchController {
	branchController := &BranchController{
		gimlet.NewController[models.BranchModel](),
	}
	branchController.ControllerHandler = branchController

	return branchController
}

func (controller *BranchController) RegisterRoutes(router gimlet.Router) {
	router.Group("/Branch/", func(r gimlet.Router) {
		r.GET("", controller.GetAll).Use(middlewares.AuthenticateAdmin)
		r.GET("{id}", controller.Get).Use(middlewares.AuthenticateAdmin)
		r.PUT("{id}", controller.Update).Use(middlewares.AuthenticateAdmin)
		r.DELETE("{id}", controller.Delete).Use(middlewares.AuthenticateAdmin)
		r.POST("", controller.Create).Use(middlewares.AuthenticateAdmin)
	})
}

func (controller *BranchController) GetService() gimlet.ServiceInterface[models.BranchModel] {
	return services.NewBranchService()
}
