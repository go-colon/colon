/**
 * @author:cb <fastopencn@gmail.com>
 * @date:2022/5/26
 **/
package auth

import (
	"github.com/go-colon/colon/app/provider/user"

	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/gin"
)

// AuthMiddleware 登录中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		envService := c.MustMake(contract.EnvKey).(contract.Env)
		userService := c.MustMake(user.UserKey).(user.Service)
		// 如果在调试模式下，根据参数的user_id 获取信息
		if envService.AppEnv() == contract.EnvDevelopment {
			userID, exist := c.DefaultQueryInt64("user_id", 0)
			if exist {
				authUser, _ := userService.GetUser(c, userID)
				if authUser != nil {
					c.Set("auth_user", authUser)
					c.Next()
					return
				}
			}
		}

		token, err := c.Cookie("colon")
		if err != nil || token == "" {
			c.ISetStatus(401).IText("请登录后操作")
			return
		}

		authUser, err := userService.VerifyLogin(c, token)
		if err != nil || authUser == nil {
			c.ISetStatus(401).IText("请登录后操作")
			return
		}

		c.Set("auth_user", authUser)

		c.Next()
	}
}

// GetAuthUser 获取已经验证的用户
func GetAuthUser(c *gin.Context) *user.User {
	t, exist := c.Get("auth_user")
	if !exist {
		return nil
	}
	return t.(*user.User)
}
