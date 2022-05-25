package app

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// ColonAppProvider 提供App的具体实现方法
type ColonAppProvider struct {
	BaseFolder string
}

// Register 注册ColonApp方法
func (h *ColonAppProvider) Register(container core.Container) core.NewInstance {
	return NewColonApp
}

// Boot 启动调用
func (h *ColonAppProvider) Boot(container core.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *ColonAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *ColonAppProvider) Params(container core.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *ColonAppProvider) Name() string {
	return contract.AppKey
}
