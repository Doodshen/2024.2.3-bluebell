/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-25 12:30:44
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-25 12:44:43
 * @FilePath: \2024.2.3 bluebell\controller\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"
	"web_app/logic"

	"github.com/gin-gonic/gin"
)

// signupHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1 获取参数和参数校验 ---controller层处理
	//2 业务处理          ---logic层
	logic.SignUp()
	//3 返回响应
	c.JSON(http.StatusOK, "ok")
}
