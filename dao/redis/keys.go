package redis

//redis key
//redis kye 使用命名空间方式，方便查询和拆分
const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"  //Zset:帖子及发帖时间
	KeyPostScoreZSet   = "post:score" //Zset 帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted" //Zset 记录用户及投票类型 ;参数是post id
)

//给redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
