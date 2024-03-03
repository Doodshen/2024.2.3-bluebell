/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 15:03:55
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 16:51:09
 * @FilePath: \2024.2.3 bluebell\logic\community.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 15:03:55
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 17:16:21
 * @FilePath: \2024.2.3 bluebell\logic\community.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetCommunityList()查询社区列表
func GetCommunityList() ([]*models.Community, error) {
	//查数据库，查到所有的community并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetailList 调用数据库查询详细信息
func GetCommunityDetailList(id int64) (*models.CommunityDetial, error) {
	return mysql.GetCommunityDetailByID(id)
}
