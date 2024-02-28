/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-28 16:30:52
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 18:54:50
 * @FilePath: \2024.2.3 bluebell\middleware\auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"strings"
	"web_app/controller"
	"web_app/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey = "userID"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		//按照空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenstring，使用之前定义好的解析jwt的函数来解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		//将当前请求的username信息保存到请求的上下文中
		c.Set(CtxUserIDKey, mc.UserID)
		c.Next() //后续的处理函数可以通过c.Get（）就可以来获取到当前登录的用户信息
	}
}
