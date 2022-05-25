package kernel

import (
	"net/http"

	"github.com/go-colon/colon/core/gin"
)

// 引擎服务
type ColonKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewColonKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &ColonKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *ColonKernelService) HttpEngine() http.Handler {
	return s.engine
}
