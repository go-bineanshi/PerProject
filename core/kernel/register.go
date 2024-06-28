package kernel

import (
	"github.com/go-bineanshi/PerProject/core/provider/interfaces"
	"sync"
)

type IServiceProvider = iprovider.ServiceProvider
type Provider = func() IServiceProvider

type Register struct {
	mu sync.RWMutex

	providers []Provider
}

var defaultRegister = NewRegister()

func NewRegister() *Register {
	return &Register{
		providers: make([]Provider, 0),
	}
}

func (r *Register) AddProvider(fn func() any) *Register {
	r.mu.Lock()
	defer r.mu.Unlock()

	provider := fn()

	// 判断是否为服务提供者
	switch p := provider.(type) {
	case IServiceProvider:
		r.providers = append(r.providers, func() IServiceProvider {
			return p
		})
	}

	return r
}

func AddProvider(f func() any) *Register {
	return defaultRegister.AddProvider(f)
}

// 获取全部服务提供者
func (this *Register) GetAllProvider() []Provider {
	return this.providers
}

// 获取全部服务提供者
func GetAllProvider() []Provider {
	return defaultRegister.GetAllProvider()
}
