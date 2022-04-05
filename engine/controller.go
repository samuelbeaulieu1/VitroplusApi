package engine

import (
	"net/http"
	"strconv"
)

type IController interface {
	RegisterRoutes(router IRouter)
}

func ParseModelId(id string) (uint, error) {
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsedId), nil
}

func ParseRouteIdentifier(key string, ctx *Context) (uint, bool) {
	if ctx.GetParam(key) == "" {
		ctx.WriteJSONError(http.StatusBadRequest, NewError("L'identifiant est obligatoire"))
		return 0, false
	}
	id, err := ParseModelId(ctx.GetParam(key))
	if err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, NewError("L'identifiant est invalide"))
		return 0, false
	}

	return id, true
}
