package demo

import (
	"github.com/go-colon/colon/core"
)

type DemoProvider struct {
	core.ServiceProvider

	c core.Container
}

func (sp *DemoProvider) Name() string {
	return DemoKey
}

func (sp *DemoProvider) Register(c core.Container) core.NewInstance {
	return NewService
}

func (sp *DemoProvider) IsDefer() bool {
	return false
}

func (sp *DemoProvider) Params(c core.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *DemoProvider) Boot(c core.Container) error {
	sp.c = c
	return nil
}
