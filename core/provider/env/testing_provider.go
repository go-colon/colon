package env

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonTestingEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *ColonTestingEnvProvider) Register(c core.Container) core.NewInstance {
	return NewColonTestingEnv
}

// Boot will called when the service instantiate
func (provider *ColonTestingEnvProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonTestingEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonTestingEnvProvider) Params(c core.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *ColonTestingEnvProvider) Name() string {
	return contract.EnvKey
}
