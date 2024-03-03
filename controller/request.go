/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-28 18:56:28
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 19:11:41
 * @FilePath: \2024.2.3 bluebell\controller\request.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 将常量进行提取，解决循环引用
const (
	CtxUserIDKey = "userID"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return

}

func getPageSize(c *gin.Context) (page, size int64) {
	pagestr := c.Query("page")
	sizestr := c.Query("size")

	page, err := strconv.ParseInt(pagestr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizestr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
