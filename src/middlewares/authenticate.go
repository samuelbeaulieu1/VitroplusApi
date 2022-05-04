package middlewares

import (
	"strings"

	"github.com/samuelbeaulieu1/gimlet"
)

func AuthenticateAdmin(route *gimlet.Route, ctx *gimlet.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(authorization, "Bearer ")
	if len(splitToken) != 2 {
		// authenticationError(route, ctx)
		return
	}
	/* token := splitToken[1]
	if payload, ok := verifySession(token); ok {
		ctx.Authentication = payload
		return
	}
	authenticationError(route, ctx) */
}

/* func verifySession(token string) (*engine.AuthTokenPayload, bool) {
	authService := services.NewAuthService()
	payload, err := authService.ValidateToken(token)
	if err != nil {
		return nil, false
	}

	valid := validateSession(payload)
	if !valid {
		return nil, false
	}
	return payload, true
}

func authenticationError(route *engine.Route, ctx *engine.Context) {
	ctx.WriteJSONError(http.StatusUnauthorized, engine.NewError("Authentification invalide"))
	route.CancelExecution()
}

func validateSession(token *engine.AuthTokenPayload) bool {
	session := entities.NewSession()
	query := &models.SessionModel{}
	id, err := utils.ParseModelId(token.Id)
	if err != nil {
		return false
	}
	query.Model = gorm.Model{
		ID: id,
	}
	res, err := session.Get(query)

	return err == nil && res.Valid
} */
