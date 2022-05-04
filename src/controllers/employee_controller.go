package controllers

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/vitroplus-api/src/middlewares"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
	"github.com/samuelbeaulieu1/vitroplus-api/src/services"
)

type EmployeeController struct {
	*gimlet.Controller[models.EmployeeModel]
}

func NewEmployeeController() *EmployeeController {
	employeeController := &EmployeeController{
		gimlet.NewController[models.EmployeeModel](),
	}
	employeeController.ControllerHandler = employeeController

	return employeeController
}

func (controller *EmployeeController) RegisterRoutes(router gimlet.Router) {
	router.Group("/Employee", func(r gimlet.Router) {
		r.POST("", controller.Create).Use(middlewares.AuthenticateAdmin)
		r.GET("", controller.GetAll).Use(middlewares.AuthenticateAdmin)
		r.GET("/DailyReport/{pin}", controller.getDailyReport)
		r.GET("/{id}", controller.Get).Use(middlewares.AuthenticateAdmin)
		r.PUT("/{id}", controller.Update).Use(middlewares.AuthenticateAdmin)
		r.DELETE("/{id}", controller.Delete).Use(middlewares.AuthenticateAdmin)
	})
}

func (controller *EmployeeController) GetService() gimlet.ServiceInterface[models.EmployeeModel] {
	return services.NewEmployeeService()
}

func (controller *EmployeeController) getDailyReport(ctx *gimlet.Context) {

}
