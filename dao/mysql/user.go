/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-25 12:36:32
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-27 21:55:29
 * @FilePath: \2024.2.3 bluebell\dao\mysql\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

//把每一步数据哭操作封装成函数
//等待logic层根据业务需求调用

const secret = "kingshen"

// CheckUserExists 检查指定数据库是否存在该用户
func CheckUserExists(username string) (err error) {
	sqlstr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user models.User) (err error) {
	//对密码进行加密
	password := encryptPassword(user.Password)
	//执行SQl入库
	sqlstr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlstr, user.UserID, user.Username, password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login登录函数
func Login(u *models.User) (err error) {
	oPassword := u.Password
	sqlstr := "select user_id,username,password from user where username=?"
	err = db.Get(u, sqlstr, u.Username)
	//判断用户存不存在
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	//查询数据库失败
	if err != nil {
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != u.Password {
		return errors.New("密码错误")
	}
	return
}
