package distributed

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// LocalDistributedProvider 提供App的具体实现方法
type LocalDistributedProvider struct {
}

// Register 注册ColonApp方法
func (h *LocalDistributedProvider) Register(container core.Container) core.NewInstance {
	return NewLocalDistributedService
}

// Boot 启动调用
func (h *LocalDistributedProvider) Boot(container core.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *LocalDistributedProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *LocalDistributedProvider) Params(container core.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}
