package id

import (
	"testing"

	tests "github.com/go-colon/colon/test"

	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/provider/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test colon console log normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&config.ColonConfigProvider{})

		err := c.Bind(&ColonIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		So(xid, ShouldNotBeEmpty)
	})
}
