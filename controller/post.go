/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 19:36:21
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 20:36:13
 * @FilePath: \2024.2.3 bluebell\controller\post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 处理创建文章post请求
func CreatePostHandler(c *gin.Context) {
	//1. 获取请求参数
	//1.1 创建模型存储参数
	post := new(models.Post)
	if err := c.ShouldBind(post); err != nil { //ShouldBindJson（） 调用内部的validator 对应到一个binding tag中进行绑定
		zap.L().Debug("测送错误", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//1.2 author_id 通过当前登录用户获取
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorID = userID

	//2.创建帖子
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic CreatePost(p) failed ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, nil)
}
