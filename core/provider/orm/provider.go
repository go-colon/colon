package orm

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// GormProvider 提供App的具体实现方法
type GormProvider struct {
}

// Register 注册方法
func (h *GormProvider) Register(container core.Container) core.NewInstance {
	return NewColonGorm
}

// Boot 启动调用
func (h *GormProvider) Boot(container core.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *GormProvider) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (h *GormProvider) Params(container core.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *GormProvider) Name() string {
	return contract.ORMKey
}
