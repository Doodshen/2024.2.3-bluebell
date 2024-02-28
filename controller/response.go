/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-28 13:15:55
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 14:30:28
 * @FilePath: \2024.2.3 bluebell\controller\response.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义返回响应的结构体
type ResponseData struct {
	Code Rescode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseError()响应错误信息
func ResponseError(c *gin.Context, code Rescode) {
	responsedata := &ResponseData{
		Code: code,
		Msg:  code.Msg(), //根据错误码去查询错误信息
		Data: nil,
	}
	c.JSON(http.StatusOK, responsedata)
}

// ResponseSuccess()响应成功信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	responsedata := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(), //根据错误码去查询错误信息
		Data: data,
	}
	c.JSON(http.StatusOK, responsedata)
}

// ResponseErrorWithMsg()响应自定义错误信息
func ResponseErrorWithMsg(c *gin.Context, code Rescode, msg interface{}) {
	responsedata := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, responsedata)
}
