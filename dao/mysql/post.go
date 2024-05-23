/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 20:06:49
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 19:50:50
 * @FilePath: \2024.2.3 bluebell\dao\mysql\post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql

import (
	"strings"
	"web_app/models"

	"github.com/jmoiron/sqlx"
)

// CreatePost将帖子加入到数据库
func CreatePost(p *models.Post) (err error) {
	//构建sql语句
	str := `insert into post(
			post_id,title,content,author_id,community_id)
			values(?,?,?,?,?)
	`
	_, err = db.Exec(str, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

// GetPostById 根据id查询单个帖子数据
func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`

	err = db.Get(post, sqlStr, id)
	return
}

// GetPostList 查找所有贴子
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlstr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	order by create_time 
	DESC
	limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlstr, (page-1)*size, size)
	return
}

// 根据给定的id列表查询帖子
func GetPostListByIDs(ids []string) (postlist []*models.Post, err error) {
	sqlstr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`

	query, args, err := sqlx.In(sqlstr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postlist, query, args...)
	return
}
