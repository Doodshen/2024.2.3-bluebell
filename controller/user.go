/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-25 12:30:44
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 15:48:48
 * @FilePath: \2024.2.3 bluebell\controller\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// signupHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {

	//1 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("SingUp with invalid param ", zap.Error(err))

		//判断err是不是validator.validationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SingUp with failed ", zap.Error(err))
		if errors.Is(err, mysql.ErrorIsNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	//1 参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))

		//判断err是不是validator.validationError类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2 业务处理
	token, err := logic.Login(p) //接收logic的token
	if err != nil {
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//3 返回响应
	ResponseSuccess(c, token)
}
