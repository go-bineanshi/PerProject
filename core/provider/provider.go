package provider

import (
	"github.com/go-bineanshi/PerProject/core/app/interfaces"
	"github.com/go-bineanshi/PerProject/core/router"
)

type ServiceProvider struct {
	App   iapp.App
	Route *router.Engine
}

func (sp *ServiceProvider) WithApp(a any) {
	sp.App = a.(iapp.App)
}
func (sp *ServiceProvider) GetApp() iapp.App {
	return sp.App
}

func (sp *ServiceProvider) WithRoute(engine *router.Engine) {
	sp.Route = engine
}

func (sp *ServiceProvider) AddRoute(fn func(*router.Engine)) {
	if sp.Route != nil {
		fn(sp.Route)
	}
}

func (sp *ServiceProvider) GetRoute() *router.Engine {
	return sp.Route
}

func (sp *ServiceProvider) Boot() {
}

func (sp *ServiceProvider) Register() {
}
