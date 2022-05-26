package user

import (
	"github.com/go-colon/colon/core"
)

type UserProvider struct {
	core.ServiceProvider

	c core.Container
}

func (sp *UserProvider) Name() string {
	return UserKey
}

func (sp *UserProvider) Register(c core.Container) core.NewInstance {
	return NewUserService
}

func (sp *UserProvider) IsDefer() bool {
	return true
}

func (sp *UserProvider) Params(c core.Container) []interface{} {
	return []interface{}{c}
}

func (sp *UserProvider) Boot(c core.Container) error {
	return nil
}

