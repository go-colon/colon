package ssh

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// SSHProvider 提供App的具体实现方法
type SSHProvider struct {
}

// Register 注册方法
func (h *SSHProvider) Register(container core.Container) core.NewInstance {
	return NewColonSSH
}

// Boot 启动调用
func (h *SSHProvider) Boot(container core.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *SSHProvider) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (h *SSHProvider) Params(container core.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *SSHProvider) Name() string {
	return contract.SSHKey
}
