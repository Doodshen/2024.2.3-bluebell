package controller

import (
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	//2 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	//1 参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {

	}

	//2 业务处理
	if err := logic.Login(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	//3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
