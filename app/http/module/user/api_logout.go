/**
 * @author:cb <fastopencn@gmail.com>
 * @date:2022/5/26
 **/
package user

import (
	"github.com/go-colon/colon/app/http/middleware/auth"
	provider "github.com/go-colon/colon/app/provider/user"
	"github.com/go-colon/colon/core/gin"
)

// Logout 用户退出
// @Summary 用户登出
// @Description 注销用户登录信息
// @Accept  json
// @Produce  json
// @Tags user
// @Success 200 {string} Message "用户退出成功"
// @Router /user/logout [get]
func (api *UserApi) Logout(c *gin.Context) {
	authUser := auth.GetAuthUser(c)
	if authUser == nil {
		c.ISetStatus(500).IText("用户未登录")
		return
	}

	userService := c.MustMake(provider.UserKey).(provider.Service)
	if err := userService.Logout(c, authUser); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	//c.ISetOkStatus().IText("用户退出成功")
	c.IRedirect("/#/login")
	return
}
