package config

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type FakeConfigProvider struct {
	FileName string
	Content  []byte
}

// Register registe a new function for make a service instance
func (provider *FakeConfigProvider) Register(c core.Container) core.NewInstance {
	return NewFakeConfig
}

// Boot will called when the service instantiate
func (provider *FakeConfigProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *FakeConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *FakeConfigProvider) Params(c core.Container) []interface{} {
	return []interface{}{provider.FileName, provider.Content}
}

// Name define the name for this service
func (provider *FakeConfigProvider) Name() string {
	return contract.ConfigKey
}
