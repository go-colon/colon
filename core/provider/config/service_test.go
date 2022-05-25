package config

import (
	"path/filepath"
	"testing"

	"github.com/go-colon/colon/core/contract"
	tests "github.com/go-colon/colon/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColonConfig_GetInt(t *testing.T) {
	container := tests.InitBaseContainer()

	Convey("test colon env normal case", t, func() {
		appService := container.MustMake(contract.AppKey).(contract.App)
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		folder := filepath.Join(appService.ConfigFolder(), envService.AppEnv())

		serv, err := NewColonConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*ColonConfig)
		timeout := conf.GetString("database.default.timeout")
		So(timeout, ShouldEqual, "10s")
	})
}
