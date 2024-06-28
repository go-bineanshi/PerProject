package route

import "github.com/go-bineanshi/PerProject/core/router"

func AddRoute(fn func(rg *router.RouterGroup)) {
	engine := router.NewRoute().Get()
	r := engine.Group("/")
	fn(r)
}
