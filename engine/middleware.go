package engine

type Middleware func(route *Route, ctx *Context)
