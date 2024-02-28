/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-28 13:15:55
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 13:43:01
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
