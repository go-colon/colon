package id

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *ColonIDProvider) Register(c core.Container) core.NewInstance {
	return NewColonIDService
}

// Boot will called when the service instantiate
func (provider *ColonIDProvider) Boot(c core.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *ColonIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *ColonIDProvider) Params(c core.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *ColonIDProvider) Name() string {
	return contract.IDKey
}
