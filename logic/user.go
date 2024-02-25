package logic

import (
	"web_app/dao/mysql"
	sf "web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp() {

	//1 判断用户存不存在
	mysql.QueryUserByUsername()
	//2 生成UID
	sf.GenID()
	//3 保存到数据库
	mysql.InserUser()
}
