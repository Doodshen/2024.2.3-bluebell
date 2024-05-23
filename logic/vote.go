package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

//投票功能
//1 用户投票数据

//简化版本投票分数
//投一票就加上432分  86400/200 -->需要200赞成票可以让你帖子续一天
/*投票情况：
direction =1 时，两种情况：
    1 之前没有投过票，现在投赞成票
	2 之前投反对票，现在投占城票
direction =0时，两种情况：
    1 之前投过赞成票，现在要取消投票
	2 之前通过返回票，现在要取消投票
direction=-1 时:有两种情况
    1.之前没有投过票，现在投反对票
	2 之前投过票，现在改投反对票

投票限制：
  每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票
  1.到期之后将redis中保存的赞成票以及反对票存储到mysql中
  2 到期之后删除KeyPostVotedZSetPF
*/

//VoteForPost 为帖子投票

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("voteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("diection", p.Direction))

	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
