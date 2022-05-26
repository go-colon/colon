/**
 * @author:cb <fastopencn@gmail.com>
 * @date:2022/5/26
 **/
package user

import (
	"github.com/go-colon/colon/app/http/middleware/auth"
	"github.com/go-colon/colon/app/provider/user"
	"github.com/go-colon/colon/core/gin"
)

type UserApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}
	if !r.IsBind(user.UserKey) {
		r.Bind(&user.UserProvider{})
	}

	// 登录
	r.POST("/user/login", api.Login)
	// 登出
	r.GET("/user/logout", auth.AuthMiddleware(), api.Logout)
	// 注册
	r.POST("/user/register", api.Register)
	// 注册验证
	r.GET("/user/register/verify", api.Verify)

	return nil
}
