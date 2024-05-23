package redis

import "web_app/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id
	//根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2 确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	//3 ZRevRange 获取有序集合中指定索引范围内的元素
	return client.ZRevRange(key, start, end).Result()
}
