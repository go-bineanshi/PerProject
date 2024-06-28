package bootstrap

import (
	"github.com/go-bineanshi/PerProject/core/kernel"
)

func Execute() {
	// 服务提供者
	providers := kernel.GetAllProvider()

	kernel.New().
		WithServiceProviders(providers).
		RunServer()
}
