package app

import (
	"github.com/go-bineanshi/PerProject/core/provider/interfaces"
	"github.com/go-bineanshi/PerProject/core/router"
	"reflect"
	"sync"
)

type App struct {
	run bool

	// 锁
	mut sync.RWMutex

	// 服务提供者
	serviceProviders []iprovider.ServiceProvider

	// 已经加载的服务提供者
	loadedProviders map[string]bool

	route *router.Engine
}

func New() *App {
	return &App{
		run:              false,
		serviceProviders: make([]iprovider.ServiceProvider, 0),
		loadedProviders:  make(map[string]bool),
	}
}

func (app *App) loadServiceProvider() {
	usedServiceProviders := make([]iprovider.ServiceProvider, 0)
	for _, p := range app.serviceProviders {
		p.WithApp(app)
		p.WithRoute(app.route)
		p.Register()
		usedServiceProviders = append(usedServiceProviders, p)
	}

	for _, sp := range usedServiceProviders {
		app.BootService(sp)
	}
}
func (app *App) BootService(s iprovider.ServiceProvider) {
	// 启动
	s.Boot()
}
func (this *App) Run() {
	this.runApp()
}

func (app *App) runApp() {
	var r *router.Engine
	r = router.New()
	router.NewRoute().With(r)

	app.route = r
	app.run = true

	app.loadServiceProvider()
	app.serverRun()
}

func (a *App) serverRun() {
	if err := a.route.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Register(f func() iprovider.ServiceProvider) iprovider.ServiceProvider {
	provide := f()
	if serviceProvider := a.GetRegister(provide); serviceProvider != nil {
		return serviceProvider
	}
	a.markAsRegistered(provide)

	if a.run {
		// 绑定 App 结构体
		provide.WithApp(a)

		// 路由
		provide.WithRoute(a.route)

		// 注册
		provide.Register()

		// 引导
		a.BootService(provide)
	}
	return provide
}

func (a *App) Registers(provides []func() iprovider.ServiceProvider) {
	for _, provide := range provides {
		a.Register(provide)
	}
}

// 注册
func (a *App) markAsRegistered(provider iprovider.ServiceProvider) {
	a.mut.Lock()
	defer a.mut.Unlock()

	a.serviceProviders = append(a.serviceProviders, provider)

	a.loadedProviders[a.GetProviderName(provider)] = true
}

func (a *App) GetLoadedProviders() map[string]bool {
	return a.loadedProviders
}

// GetProviderName 反射获取服务提供者名称
func (a *App) GetProviderName(provider any) (name string) {
	p := reflect.TypeOf(provider)

	if p.Kind() == reflect.Pointer {
		p = p.Elem()
		name = "*"
	}

	pkgPath := p.PkgPath()

	if pkgPath != "" {
		name += pkgPath + "."
	}

	return name + p.Name()
}

func (a *App) GetRegister(provide any) iprovider.ServiceProvider {
	var name string
	switch t := provide.(type) {
	case iprovider.ServiceProvider:
		name = a.GetProviderName(t)
	case string:
		name = t
	}
	if name != "" {
		for _, serviceProvider := range a.serviceProviders {
			if a.GetProviderName(serviceProvider) == name {
				return serviceProvider
			}
		}
	}
	return nil
}

func (this *App) ProviderIsLoaded(provider string) bool {
	this.mut.RLock()
	defer this.mut.RUnlock()

	if _, ok := this.loadedProviders[provider]; ok {
		return true
	}

	return false
}
