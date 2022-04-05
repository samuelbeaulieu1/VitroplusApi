package main

import (
	"github.com/samuelbeaulieu1/vitroplus-api/engine"
	"github.com/samuelbeaulieu1/vitroplus-api/src/middlewares"
)

func main() {
	instance := engine.NewEngine()

	instance.LoadConfig("./config.json")
	instance.Use(middlewares.CORS)

	registerControllers(instance)
	instance.Run()
}

func registerControllers(instance *engine.Engine) {
	controllers := []engine.IController{}

	for _, controller := range controllers {
		controller.RegisterRoutes(instance)
	}
}
