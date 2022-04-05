package middlewares

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/samuelbeaulieu1/vitroplus-api/server/engine"
)

func CORS(route *engine.Route, ctx *engine.Context) {
	method := ctx.Request.Method

	allowOrigin(ctx)
	allowCredentials(ctx)

	if method == http.MethodOptions {
		allowHeaders(ctx)
		allowMethods(ctx)
		ctx.Status(http.StatusOK)
		route.CancelExecution()
	}
}

func allowOrigin(ctx *engine.Context) {
	origin := ctx.Request.Header.Get("Origin")

	for _, domain := range ctx.Engine.AllowOrigin {
		if match, _ := regexp.Match(domain, []byte(origin)); match || domain == origin || domain == "*" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
	}
}

func allowMethods(ctx *engine.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(ctx.Engine.AllowMethods, ", "))
}

func allowHeaders(ctx *engine.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(ctx.Engine.AllowHeaders, ", "))
}

func allowCredentials(ctx *engine.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(ctx.Engine.AllowCredentials))
}
