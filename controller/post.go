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
	"strconv"
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

// GetPostDetailHandler 根据ID查询帖子详情
func GetPostyDetailHandler(c *gin.Context) {
	//1 参数校验 --get请求就是查看url上的参数
	pidstr := c.Param("id")                      //查询出来的pid是string
	pid, err := strconv.ParseInt(pidstr, 10, 64) //10 为十进制，64表示int64
	if err != nil {
		zap.L().Error("get post detail invalid params ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2 查询帖子详情
	data, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("get post detaile failed ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	//参数校验，获取页数和条数
	page, size := getPageSize(c)
	//查询数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("GetPostlist failed:%v:", zap.Error(err))
	}

	//返回响应
	ResponseSuccess(c, data)

}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}
