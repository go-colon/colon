package log

import (
	"os"

	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/provider/log/formatter"
	"github.com/go-colon/colon/core/provider/log/services"
)

// ColonTestingLog 是 Env 的具体实现
type ColonTestingLog struct {
}

// NewColonTestingLog 测试日志，直接打印到控制台
func NewColonTestingLog(params ...interface{}) (interface{}, error) {
	log := &services.ColonConsoleLog{}

	log.SetLevel(contract.DebugLevel)
	log.SetCtxFielder(nil)
	log.SetFormatter(formatter.TextFormatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	return log, nil
}
