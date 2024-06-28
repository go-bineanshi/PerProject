package bootstrap

import (
	"github.com/go-bineanshi/PerProject/apps/demo/provider"
	"github.com/go-bineanshi/PerProject/core/kernel"
)

func Boot() {
	kernel.AddProvider(func() any {
		return &provider.DemoProvider{}
	})
}
