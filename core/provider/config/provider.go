package config

import (
	"path/filepath"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonConfigProvider struct{}

// Register registe a new function for make a service instance
func (provider *ColonConfigProvider) Register(c core.Container) core.NewInstance {
	return NewColonConfig
}

// Boot will called when the service instantiate
func (provider *ColonConfigProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonConfigProvider) Params(c core.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

/// Name define the name for this service
func (provider *ColonConfigProvider) Name() string {
	return contract.ConfigKey
}
