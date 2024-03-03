/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 20:06:49
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 19:50:50
 * @FilePath: \2024.2.3 bluebell\dao\mysql\post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql

import "web_app/models"

//CreatePost将帖子加入到数据库
func CreatePost(p *models.Post) (err error) {
	//构建sql语句
	str := `insert into post(
			post_id,title,content,author_id,community_id)
			values(?,?,?,?,?)
	`
	_, err = db.Exec(str, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

//GetPostByID查询帖子通过帖子ID
func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`

	err = db.Get(post, sqlStr, id)
	return
}

//GetPostList 查找所有贴子
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlstr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlstr, (page-1)*size, size)
	return
}
