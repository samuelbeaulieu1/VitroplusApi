package controllers

import (
	"net/http"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
	"github.com/samuelbeaulieu1/vitroplus-api/src/middlewares"
	"github.com/samuelbeaulieu1/vitroplus-api/src/services"
)

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (controller *AdminController) RegisterRoutes(router gimlet.Router) {
	router.Group("/Admin/", func(r gimlet.Router) {
		r.PUT("", controller.Update).Use(middlewares.AuthenticateAdmin)
		r.POST("", controller.Auth)
	})
}

func (controller *AdminController) Update(ctx *gimlet.Context) {
	var req classes.UpdateAdminRequest
	ctx.ParseBody(&req)

	if err := services.NewAdminService().UpdatePassword(&req); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		ctx.WriteJSONResponse(&responses.RequestResponseMessage{
			Message: "Le mot de passe a été modifié",
		})
	}
}

func (controller *AdminController) Auth(ctx *gimlet.Context) {
	var req classes.AdminAuthRequest
	ctx.ParseBody(&req)

	if session, err := services.NewAdminService().CreateSession(&req); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		ctx.WriteJSONResponse(session)
	}
}
