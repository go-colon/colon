package util

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-colon/colon/core/gin"
	"github.com/go-colon/colon/core/provider/log"
	"github.com/go-colon/colon/core/util/goroutine"
	tests "github.com/go-colon/colon/test"
)

func TestSafeGo(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&log.ColonTestingLogProvider{})

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	goroutine.SafeGo(ctx, func() {
		time.Sleep(1 * time.Second)
		return
	})
	t.Log("safe go main start")
	time.Sleep(2 * time.Second)
	t.Log("safe go main end")

	goroutine.SafeGo(ctx, func() {
		time.Sleep(1 * time.Second)
		panic("safe go test panic")
	})
	t.Log("safe go2 main start")
	time.Sleep(2 * time.Second)
	t.Log("safe go2 main end")

}

func TestSafeGoAndWait(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&log.ColonTestingLogProvider{})

	errStr := "safe go test error"
	t.Log("safe go and wait start", time.Now().String())
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err := goroutine.SafeGoAndWait(ctx, func() error {
		time.Sleep(1 * time.Second)
		return errors.New(errStr)
	}, func() error {
		time.Sleep(2 * time.Second)
		return nil
	}, func() error {
		time.Sleep(3 * time.Second)
		return nil
	})
	t.Log("safe go and wait end", time.Now().String())

	if err == nil {
		t.Error("err not be nil")
	} else if err.Error() != errStr {
		t.Error("err content not same")
	}

	// panic error
	err = goroutine.SafeGoAndWait(ctx, func() error {
		time.Sleep(1 * time.Second)
		return errors.New(errStr)
	}, func() error {
		time.Sleep(2 * time.Second)
		panic("test2")
	}, func() error {
		time.Sleep(3 * time.Second)
		return nil
	})
	if err == nil {
		t.Error("err not be nil")
	} else if err.Error() != errStr {
		t.Error("err content not same")
	}
}
