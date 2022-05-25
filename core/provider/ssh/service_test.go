package ssh

import (
	"testing"

	"github.com/go-colon/colon/core/provider/config"
	"github.com/go-colon/colon/core/provider/log"
	tests "github.com/go-colon/colon/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestColonSSHService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.ColonConfigProvider{})
	container.Bind(&log.ColonLogServiceProvider{})

	Convey("test get client", t, func() {
		colonRedis, err := NewColonSSH(container)
		So(err, ShouldBeNil)
		service, ok := colonRedis.(*ColonSSH)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
