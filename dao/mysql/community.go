/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 15:08:49
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 17:14:44
 * @FilePath: \2024.2.3 bluebell\dao\mysql\community.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

// GetCommunityList()查询数据库中社区列表
func GetCommunityList() (CommunityList []*models.Community, err error) {

	sqlstr := "select community_id,community_name from community"

	if err := db.Select(&CommunityList, sqlstr); err != nil { //自动将数据拷贝到这个结构体了
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}

	}
	return
}

// GetCommunityDetailList  根据id查询数据库中列表详情
func GetCommunityDetailList(id int64) (community *models.CommunityDetialList, err error) {
	community = new(models.CommunityDetialList)

	str := `select community_id,community_name,introduction,create_time from community where community_id = ? `
	if err = db.Get(community, str, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
