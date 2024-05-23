package controller

import (
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//投票

type VoteData struct {
	//UserID
	PostID    int64 `json:"post_id"`          //帖子ID
	Direction int   `json:"direction,string"` //1:赞 2:踩

}

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译去除错误提示信息中的结构体  //前端不需要知道结构体
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	//获取当前请求的用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//具体投票逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, nil)
}
