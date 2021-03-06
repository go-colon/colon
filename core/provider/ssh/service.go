package ssh

import (
	"context"
	"sync"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
	"golang.org/x/crypto/ssh"
)

// ColonSSH colon框架的ssh实现
type ColonSSH struct {
	container core.Container         // 服务容器
	clients   map[string]*ssh.Client // key为uniqKey, value为ssh.Client(连接池）

	lock *sync.RWMutex
}

// NewColonSSH 实例化Client
func NewColonSSH(params ...interface{}) (interface{}, error) {
	container := params[0].(core.Container)
	clients := make(map[string]*ssh.Client)
	lock := &sync.RWMutex{}
	return &ColonSSH{
		container: container,
		clients:   clients,
		lock:      lock,
	}, nil
}

// GetClient 获取Client实例
func (app *ColonSSH) GetClient(option ...contract.SSHOption) (*ssh.Client, error) {
	logService := app.container.MustMake(contract.LogKey).(contract.Log)
	// 读取默认配置
	config := GetBaseConfig(app.container)

	// option对opt进行修改
	for _, opt := range option {
		if err := opt(app.container, config); err != nil {
			return nil, err
		}
	}

	// 如果最终的config没有设置dsn,就生成dsn
	key := config.UniqKey()

	// 判断是否已经实例化了
	app.lock.RLock()
	if db, ok := app.clients[key]; ok {
		app.lock.RUnlock()
		return db, nil
	}
	app.lock.RUnlock()

	// 没有实例化,那么就要进行实例化操作
	app.lock.Lock()
	defer app.lock.Unlock()

	// 实例化
	addr := config.Host + ":" + config.Port
	client, err := ssh.Dial(config.NetWork, addr, config.ClientConfig)
	if err != nil {
		logService.Error(context.Background(), "ssh dial error", map[string]interface{}{
			"err":  err,
			"addr": addr,
		})
	}

	// 挂载到map中，结束配置
	app.clients[key] = client

	return client, nil
}
