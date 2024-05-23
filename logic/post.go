package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	sf "web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost() 创建帖子，放入数据库
func CreatePost(p *models.Post) (err error) {
	//1. 创建帖子,将模型中的各个数据与数据库进行对齐
	//1.1文章记录id，雪花算法
	p.ID = sf.GenID()

	//2 执行入库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return err

}

func GetPostDetail(id int64) (data *models.ApiPostDetail, err error) {
	//查询数据组合接口想要的数据
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql GetPostByPostID failed ", zap.Int64("id", id), zap.Error(err))
		return
	}
	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql GetUserByID failed ", zap.Int64("Post.AuthorID", post.AuthorID), zap.Error(err))
		//如果少了这个return 一旦出现错误这个user就会是空指针下面同理
		return
	}

	//根据社区id查询社区详情
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql GetCommunityDetailList", zap.Error(err))
		return
	}

	//拼装数据 防止空指针引用
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetial: community,
	}
	return data, err
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	//查询数据
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	//构建数据
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql GetUserByID failed ", zap.Int64("Post.AuthorID", post.AuthorID), zap.Error(err))
			//如果少了这个return 一旦出现错误这个user就会是空指针下面同理
			continue
		}

		//根据社区id查询社区详情
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql GetCommunityDetailList", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetial: community,
		}
		data = append(data, postDetail)

	}
	return
}
