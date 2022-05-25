package log

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonTestingLogProvider struct {
}

// Register registe a new function for make a service instance
func (provider *ColonTestingLogProvider) Register(c core.Container) core.NewInstance {
	return NewColonTestingLog
}

// Boot will called when the service instantiate
func (provider *ColonTestingLogProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonTestingLogProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonTestingLogProvider) Params(c core.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *ColonTestingLogProvider) Name() string {
	return contract.LogKey
}
