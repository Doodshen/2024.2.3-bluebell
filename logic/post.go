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

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//2 从redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data ")
		return
	}
	//3 根据id去数据库查询帖子详细信息
	//返回的数据还要按照给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//4 将帖子的作者及其分区信息查询出来填充到帖子中
	for idx, post := range posts {
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
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetial: community,
		}
		data = append(data, postDetail)
	}
	return

}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 2. 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetCommunityPostIDsInOrder", zap.Any("ids", ids))
	// 3. 根据id去MySQL数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetial: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostListNew  将两个查询帖子列表逻辑合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 根据请求参数的不同，执行不同的逻辑。
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return
}
