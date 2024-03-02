package mysql

import "web_app/models"

func CreatePost(p *models.Post) (err error) {
	//构建sql语句
	str := `insert into post(
			post_id,title,content,author_id,community_id)
			values(?,?,?,?,?)
	`
	_, err = db.Exec(str, p.ID, p.Title, p.Context, p.AuthorID, p.CommunityID)
	return err
}
