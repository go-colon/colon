package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-colon/colon/core/provider/config"
	"github.com/go-colon/colon/core/provider/log"
	tests "github.com/go-colon/colon/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestColonService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.ColonConfigProvider{})
	container.Bind(&log.ColonLogServiceProvider{})

	Convey("test get client", t, func() {
		colonRedis, err := NewColonRedis(container)
		So(err, ShouldBeNil)
		service, ok := colonRedis.(*ColonRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
