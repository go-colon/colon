package config

import (
	"testing"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/provider/app"
	"github.com/go-colon/colon/core/provider/env"
	tests "github.com/go-colon/colon/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColonConfig_Normal(t *testing.T) {
	Convey("test colon config normal case", t, func() {
		basePath := tests.BasePath
		c := core.NewColonContainer()
		c.Bind(&app.ColonAppProvider{BaseFolder: basePath})
		c.Bind(&env.ColonEnvProvider{})

		err := c.Bind(&ColonConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.default.host"), ShouldEqual, "localhost")
		So(conf.GetInt("database.default.port"), ShouldEqual, 3306)
		//So(conf.GetFloat64("database.default.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.default")
		So(maps, ShouldContainKey, "host")
		So(maps["host"], ShouldEqual, "localhost")

		maps2 := conf.GetStringMapString("database.default")
		So(maps2["host"], ShouldEqual, "localhost")

		type Mysql struct {
			Host string `yaml:"host"`
		}
		ms := &Mysql{}
		err = conf.Load("database.default", ms)
		So(err, ShouldBeNil)
		So(ms.Host, ShouldEqual, "localhost")
	})
}
