package iprovider

import (
	"github.com/go-bineanshi/PerProject/core/router"
)

type ServiceProvider interface {
	WithApp(any)

	WithRoute(*router.Engine)

	GetRoute() *router.Engine

	Register()

	Boot()
}
