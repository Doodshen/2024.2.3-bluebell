/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-25 12:33:58
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-28 16:22:45
 * @FilePath: \2024.2.3 bluebell\logic\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	sf "web_app/pkg/snowflake"
)

//存放业务逻辑的代码

// 注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {

	//1 判断用户存不存在
	if err := mysql.CheckUserExists(p.UserName); err != nil {
		return err
	}

	//2 生成UID
	userID := sf.GenID()
	//构建User实例
	u := models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
	}

	//3 保存到数据库
	return mysql.InsertUser(u)

}

// 登录业务逻辑
func Login(p *models.ParamLogin) (token string, err error) {

	//1 构建结构体
	user := &models.User{
		Username: p.UserName,
		Password: p.Password,
	}
	//传入数据库  传递的是一个指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	//生成jwt
	token, err = jwt.GenToken(user.UserID, user.Username)
	return token, err

}
