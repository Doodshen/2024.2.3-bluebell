package redis

//redis key
//redis kye 使用命名空间方式，方便查询和拆分
//通过zinterstore 命令，将多个zset 或者 一个set和zset 合并，并保存到新的zset中
const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"  //Zset:帖子及发帖时间
	KeyPostScoreZSet   = "post:score" //Zset 帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted" //Zset 记录用户及投票类型 ;参数是post id
	KeyCommunitySetPF  = "community:" //set  保存每个分区下帖子的id
)

//给redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
