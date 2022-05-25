package cache

import (
	"strings"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/provider/cache/services"
)

// ColonCacheProvider 服务提供者
type ColonCacheProvider struct {
	core.ServiceProvider

	Driver string // Driver
}

// Register 注册一个服务实例
func (l *ColonCacheProvider) Register(c core.Container) core.NewInstance {
	if l.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewMemoryCache
		}

		cs := tcs.(contract.Config)
		l.Driver = strings.ToLower(cs.GetString("cache.driver"))
	}

	// 根据driver的配置项确定
	switch l.Driver {
	case "redis":
		return services.NewRedisCache
	case "memory":
		return services.NewMemoryCache
	default:
		return services.NewMemoryCache
	}
}

// Boot 启动的时候注入
func (l *ColonCacheProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (l *ColonCacheProvider) IsDefer() bool {
	return true
}

// Params 定义要传递给实例化方法的参数
func (l *ColonCacheProvider) Params(c core.Container) []interface{} {
	return []interface{}{c}
}

// Name 定义对应的服务字符串凭证
func (l *ColonCacheProvider) Name() string {
	return contract.CacheKey
}
