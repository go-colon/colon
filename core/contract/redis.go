package contract

import (
	"fmt"

	"github.com/go-colon/colon/core"
	"github.com/go-redis/redis/v8"
)

const RedisKey = "colon:redis"

// RedisOption 初始化的时候的选项
type RedisOption func(container core.Container, config *RedisConfig) error

// RedisService 表示一个redis服务
type RedisService interface {
	// GetClient 获取redis连接实例
	GetClient(option ...RedisOption) (*redis.Client, error)
}

// RedisConfig 为colon定义的Redis配置结构
type RedisConfig struct {
	*redis.Options
}

// UniqKey 用来唯一标识一个RedisConfig配置
func (config *RedisConfig) UniqKey() string {
	return fmt.Sprintf("%v_%v_%v_%v", config.Addr, config.DB, config.Username, config.Network)
}
