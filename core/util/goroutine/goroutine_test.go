package main

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-colon/colon/core/gin"
	tests "github.com/go-colon/colon/test"
)

func TestSafeGo(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&log.colonTestingLogProvider{})

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
