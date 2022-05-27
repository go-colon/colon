package main

import (
	"github.com/go-colon/colon/app/console"
	"github.com/go-colon/colon/app/http"
	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/provider/app"
	"github.com/go-colon/colon/core/provider/cache"
	"github.com/go-colon/colon/core/provider/config"
	"github.com/go-colon/colon/core/provider/distributed"
	"github.com/go-colon/colon/core/provider/env"
	"github.com/go-colon/colon/core/provider/id"
	"github.com/go-colon/colon/core/provider/kernel"
	"github.com/go-colon/colon/core/provider/log"
	"github.com/go-colon/colon/core/provider/orm"
	"github.com/go-colon/colon/core/provider/redis"
	"github.com/go-colon/colon/core/provider/ssh"
	"github.com/go-colon/colon/core/provider/trace"
)

func main() {
	// 初始化服务容器
	container := core.NewColonContainer()
	// 绑定App服务提供者
	container.Bind(&app.ColonAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.ColonEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.ColonConfigProvider{})
	container.Bind(&id.ColonIDProvider{})
	container.Bind(&trace.ColonTraceProvider{})
	container.Bind(&log.ColonLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.ColonCacheProvider{})
	container.Bind(&ssh.SSHProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.ColonKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
