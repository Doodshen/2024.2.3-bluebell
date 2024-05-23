/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 19:39:26
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 16:55:36
 * @FilePath: \2024.2.3 bluebell\models\post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import "time"

type Post struct {
	//与表中的数据对应，并且采用内存对齐

	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
}

//定义帖子详情结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               //嵌入帖子结构体
	*CommunityDetial `json:"communtiy"` //嵌入贴子详情

}
