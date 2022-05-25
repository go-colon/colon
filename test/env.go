package test

import (
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/provider/app"
	"github.com/go-colon/colon/core/provider/env"
)

const (
	BasePath = "/Users/cb/go/src/yuewen/colon/"
)

func InitBaseContainer() core.Container {
	// 初始化服务容器
	container := core.NewColonContainer()
	// 绑定App服务提供者
	container.Bind(&app.ColonAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.ColonTestingEnvProvider{})
	return container
}
