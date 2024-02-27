package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	sf "web_app/pkg/snowflake"
)

//存放业务逻辑的代码

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
