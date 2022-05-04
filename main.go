package main

import (
	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/lang"
	"github.com/samuelbeaulieu1/gimlet/middlewares"
	"github.com/samuelbeaulieu1/vitroplus-api/src/controllers"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dao"
)

func main() {
	instance := gimlet.NewEngine()

	instance.LoadConfig("./config.json")
	instance.Use(middlewares.CORS)

	registerControllers(instance)
	dao.InitConnection()
	lang.Set(lang.FR)
	instance.Run()
}

func registerControllers(instance *gimlet.Engine) {
	instance.Group("/v1", func(r gimlet.Router) {
		controllers.NewClockController().RegisterRoutes(r)
		controllers.NewBranchController().RegisterRoutes(r)
		controllers.NewEmployeeController().RegisterRoutes(r)
	})
}
