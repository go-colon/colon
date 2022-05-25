package redis

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// RedisProvider 提供App的具体实现方法
type RedisProvider struct {
}

// Register 注册方法
func (h *RedisProvider) Register(container core.Container) core.NewInstance {
	return NewColonRedis
}

// Boot 启动调用
func (h *RedisProvider) Boot(container core.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *RedisProvider) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (h *RedisProvider) Params(container core.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *RedisProvider) Name() string {
	return contract.RedisKey
}
