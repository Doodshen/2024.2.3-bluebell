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

//GetPostListHandler2升级版帖子列表接口
//根据前端传来的参数（按创建时间）动态获取帖子列表
//按创建时间排序，或者 按照分数排序
//1 获取参数
//2 去redis查询id列表
//3 根据id列表去数据库查询帖子详情信息

func GetPostListHandler2(c *gin.Context) {
	//Get请求参数：/api/v1/posts2?page=1&size=10&order=time  query string参数
	//参数校验，获取页数和条数
	//使用c.shouldbindquery（） 使用gin框架的默认解析器，解析query string参数 根据请求的数据类型选择相应的方法去获取参数
	//c.shouldbindquery（） 如果请求种携带是是json格式的数据，才能用这个方法获取数据

	//初始化结构体时指定初始参数
	p := models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list failed:", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//查询数据:去redis种查询id列表
	data, err := logic.GetPostList2(&p)
	if err != nil {
		zap.L().Error("GetPostlist failed:%v:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	//返回响应
	ResponseSuccess(c, data)

}
