/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-01-30 23:52:38
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-01-30 23:53:17
 * @FilePath: \web-app2\test\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	setting "web_app/settings"
)

func main() {
	fmt.Println(setting.Conf.Port)
}
