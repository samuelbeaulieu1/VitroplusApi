package middlewares

import (
	"net/http"
	"strings"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/services"
)

func AuthenticateAdmin(route *gimlet.Route, ctx *gimlet.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(authorization, "Bearer ")
	if len(splitToken) != 2 {
		authenticationError(route, ctx)
		return
	}
	token := splitToken[1]
	if ok := verifySession(token); !ok {
		authenticationError(route, ctx)
	}
}

func verifySession(token string) bool {
	authService := services.NewAdminService()
	err := authService.ValidateToken(token)
	return err == nil
}

func authenticationError(route *gimlet.Route, ctx *gimlet.Context) {
	ctx.WriteJSONError(http.StatusUnauthorized, responses.NewError("Authentification invalide"))
	route.CancelExecution()
}
