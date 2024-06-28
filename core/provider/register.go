package provider

import (
	"github.com/go-bineanshi/PerProject/core/provider/interfaces"
	"sync"
)

// 锁
var lock = new(sync.RWMutex)

var instance *Register
var once sync.Once

type Register struct {
	providers []func() iprovider.ServiceProvider
}

func NewRegister() *Register {
	once.Do(func() {
		providers := make([]func() iprovider.ServiceProvider, 0)

		instance = &Register{
			providers: providers,
		}
	})
	return instance
}

func (r *Register) Append(provider func() iprovider.ServiceProvider) {
	lock.Lock()
	defer lock.Unlock()
	r.providers = append(r.providers, provider)
}

func (r *Register) GetAll() []func() iprovider.ServiceProvider {
	return r.providers
}

func AppendProvider(f func() iprovider.ServiceProvider) {
	NewRegister().Append(f)
}

func GetAllProvider() []func() iprovider.ServiceProvider {
	return NewRegister().GetAll()
}
