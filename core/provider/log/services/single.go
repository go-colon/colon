package services

import (
	"os"
	"path/filepath"

	"github.com/go-colon/colon/core/util"

	"github.com/go-colon/colon/core"
	"github.com/go-colon/colon/core/contract"

	"github.com/pkg/errors"
)

type ColonSingleLog struct {
	ColonLog

	folder string
	file   string
	fd     *os.File
}

// NewColonSingleLog params sequence: level, ctxFielder, Formatter, map[string]interface(folder/file)
func NewColonSingleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(core.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	appService := c.MustMake(contract.AppKey).(contract.App)
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	log := &ColonSingleLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	log.folder = folder
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	log.file = "colon.log"
	if configService.IsExist("log.file") {
		log.file = configService.GetString("log.file")
	}

	fd, err := os.OpenFile(filepath.Join(log.folder, log.file), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "open log file err")
	}

	log.SetOutput(fd)
	log.c = c

	return log, nil
}
