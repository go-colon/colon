package kernel

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/gin"
)

// ColonKernelProvider 提供web引擎
type ColonKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register 注册服务提供者
func (provider *ColonKernelProvider) Register(c core.Container) core.NewInstance {
	return NewColonKernelService
}

// Boot 启动的时候判断是否由外界注入了Engine，如果注入的化，用注入的，如果没有，重新实例化
func (provider *ColonKernelProvider) Boot(c core.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// IsDefer 引擎的初始化我们希望开始就进行初始化
func (provider *ColonKernelProvider) IsDefer() bool {
	return false
}

// Params 参数就是一个HttpEngine
func (provider *ColonKernelProvider) Params(c core.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

// Name 提供凭证
func (provider *ColonKernelProvider) Name() string {
	return contract.KernelKey
}
