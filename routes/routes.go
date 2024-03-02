/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-03 14:58:10
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 16:39:48
 * @FilePath: \2024.2.3 bluebell\routes\routes.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-01-30 21:03:01
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 11:19:54
 * @FilePath: \web-app2\routes\routes.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	"net/http"
	controller "web_app/controller"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(mode string) *gin.Engine {
	//判断项目模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置为发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAutoMiddleware()) //应用jwt中间件

	{
		v1.GET("/community", controller.CommumityHandler)
		v1.GET("/community/:id", controller.CommumityDetailHandler)
	}

	//模拟登录以后才能使用的功能
	r.GET("/ping", middleware.JWTAutoMiddleware(), func(ctx *gin.Context) {
		//登陆的用户返回成功
		ctx.String(http.StatusOK, "duiduidui")
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "没有发现路由")
	})
	return r
}
