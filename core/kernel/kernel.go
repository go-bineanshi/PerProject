package kernel

import (
	"github.com/go-bineanshi/PerProject/core/app"
	"github.com/go-bineanshi/PerProject/core/provider"
	"github.com/go-bineanshi/PerProject/core/provider/interfaces"
)

type Kernel struct {
	providers []func() iprovider.ServiceProvider
}

func New() *Kernel {
	return &Kernel{}
}

func (k *Kernel) RunServer() {
	k.runApp()
}

func (k *Kernel) runApp() {
	newApp := app.New()
	k.loadServiceProvider()

	allProviders := provider.GetAllProvider()
	newApp.Registers(allProviders)
	newApp.Run()

}

func (this *Kernel) loadServiceProvider() {
	if len(this.providers) > 0 {
		for _, p := range this.providers {
			provider.AppendProvider(p)
		}
	}
}

// 添加服务提供者
func (this *Kernel) WithServiceProvider(f func() iprovider.ServiceProvider) *Kernel {
	this.providers = append(this.providers, f)

	return this
}

// 批量添加服务提供者
func (this *Kernel) WithServiceProviders(funcs []func() iprovider.ServiceProvider) *Kernel {
	if len(funcs) > 0 {
		for _, f := range funcs {
			this.WithServiceProvider(f)
		}
	}

	return this
}
