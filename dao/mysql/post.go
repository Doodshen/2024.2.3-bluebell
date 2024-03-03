/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 20:06:49
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 16:26:38
 * @FilePath: \2024.2.3 bluebell\dao\mysql\post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql

import "web_app/models"

func CreatePost(p *models.Post) (err error) {
	//构建sql语句
	str := `insert into post(
			post_id,title,content,author_id,community_id)
			values(?,?,?,?,?)
	`
	_, err = db.Exec(str, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`

	err = db.Get(post, sqlStr, id)
	return
}
