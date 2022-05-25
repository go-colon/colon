package env

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *ColonEnvProvider) Register(c core.Container) core.NewInstance {
	return NewColonEnv
}

// Boot will called when the service instantiate
func (provider *ColonEnvProvider) Boot(c core.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonEnvProvider) Params(c core.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *ColonEnvProvider) Name() string {
	return contract.EnvKey
}
