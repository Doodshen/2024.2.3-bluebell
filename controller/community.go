/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 14:58:29
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 16:10:01
 * @FilePath: \2024.2.3 bluebell\controller\community.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommumityHandler(c *gin.Context) {
	//路由处理，并响应数据
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() data err ")
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
