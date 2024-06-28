package provider

import (
	demoRoute "github.com/go-bineanshi/DemoSubItem/routes"
	"github.com/go-bineanshi/PerProject/core/provider"
	"github.com/go-bineanshi/PerProject/core/route"
	"github.com/go-bineanshi/PerProject/core/router"
)

type DemoProvider struct {
	provider.ServiceProvider
}

func (sp *DemoProvider) Boot() {
	// 启动服务
	sp.loadRoute()
}

func (sp *DemoProvider) loadRoute() {
	route.AddRoute(func(rg *router.RouterGroup) {
		demoRoute.Route(rg)
	})
}
