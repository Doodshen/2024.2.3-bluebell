/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-27 20:11:23
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 15:19:51
 * @FilePath: \2024.2.3 bluebell\models\params.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//ParamSignUp注册参数
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//ParamLogin登录参数
type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ParamPostData 投票数据
type ParamVoteData struct {
	//UserID
	PostID    string `json:"post_id" binding:"required"`                         //帖子ID
	Direction int8   `json:"direction,string" binding:"required,oneof=1 0 -1 " ` //1:赞 2:踩

}

//ParamPostList 获取帖子里欸包query string 参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
