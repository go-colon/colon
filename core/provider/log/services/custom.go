package services

import (
	"io"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"
)

type ColonCustomLog struct {
	ColonLog
}

func NewColonCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(core.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &ColonConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
