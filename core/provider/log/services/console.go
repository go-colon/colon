package services

import (
	"os"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

// ColonConsoleLog 控制台输出
type ColonConsoleLog struct {
	ColonLog
}

// NewColonConsoleLog 实例化ColonConsoleLog
func NewColonConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(core.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &ColonConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
