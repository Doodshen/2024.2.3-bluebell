/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 14:58:29
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 17:15:55
 * @FilePath: \2024.2.3 bluebell\controller\community.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommumityHandler 返回社区列表
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

// CommumityDetailHandler() 获取请求id并查询详情并返回详情
func CommumityDetailHandler(c *gin.Context) {
	//获取请求id好---但是是一个json格式需要转换为id
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	//查询所有的社区（community_id,community_name）以列表形式存在
	data, err := logic.GetCommunityDetailList(id)
	if err != nil {
		zap.L().Error("logic GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
