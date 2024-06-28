package iapp

import (
	"github.com/go-bineanshi/PerProject/core/provider/interfaces"
)

type App interface {
	// Register 注册服务提供者
	Register(func() iprovider.ServiceProvider) iprovider.ServiceProvider

	// Registers 批量注册服务提供者
	Registers([]func() iprovider.ServiceProvider)

	// GetRegister 获取
	GetRegister(any) iprovider.ServiceProvider

	// GetProviderName 反射获取服务提供者名称
	GetProviderName(any) string

	GetLoadedProviders() map[string]bool
}
