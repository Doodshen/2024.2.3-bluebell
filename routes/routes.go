/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-01-30 21:03:01
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-19 20:59:32
 * @FilePath: \web-app2\routes\routes.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	setting "web_app/settings"

	"github.com/gin-gonic/gin"
)

//注册路由

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, setting.Conf.Version)
	})

	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	return r
}
