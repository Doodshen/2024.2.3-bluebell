package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	sf "web_app/pkg/snowflake"
)

// CreatePost() 创建帖子，放入数据库
func CreatePost(p *models.Post) (err error) {
	//1. 创建帖子,将模型中的各个数据与数据库进行对齐
	//1.1文章记录id，雪花算法
	p.ID = sf.GenID()

	//2 执行入库
	return mysql.CreatePost(p)
}

func GetPostDetail(id int64) (data *models.Post, err error) {
	return mysql.GetPostByID(id)
}
