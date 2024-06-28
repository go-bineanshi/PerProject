package router

import "sync"

var instanceRoute *Route
var onceRoute sync.Once

type Route struct {
	// 路由
	routeEngine *Engine
}

func NewRoute() *Route {
	onceRoute.Do(func() {
		instanceRoute = &Route{}
	})

	return instanceRoute
}

func (r *Route) With(engine *Engine) {
	r.routeEngine = engine
}

func (r *Route) Get() *Engine {
	return r.routeEngine
}

func (r *Route) GetRoutes() RoutesInfo {
	return r.routeEngine.Routes()
}

func (r *Route) GetRouteMap() map[string]any {
	routes := r.GetRoutes()

	newRoutes := make(map[string]any)
	for _, v := range routes {
		if newRoute, ok := newRoutes[v.Method]; ok {
			newRoute = append(newRoute.([]string), v.Path)
			newRoutes[v.Method] = newRoute
		} else {
			newRoutes[v.Method] = []string{v.Path}
		}
	}

	return newRoutes
}

func (r *Route) GetLastRoute() RouteInfo {
	routes := r.routeEngine.Routes()

	return routes[len(routes)-1]
}
