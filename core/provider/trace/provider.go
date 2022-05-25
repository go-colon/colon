package trace

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonTraceProvider struct {
	c core.Container
}

// Register registe a new function for make a service instance
func (provider *ColonTraceProvider) Register(c core.Container) core.NewInstance {
	return NewColonTraceService
}

// Boot will called when the service instantiate
func (provider *ColonTraceProvider) Boot(c core.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonTraceProvider) Params(c core.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *ColonTraceProvider) Name() string {
	return contract.TraceKey
}
