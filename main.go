/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-03 14:58:10
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-27 20:59:01
 * @FilePath: \2024.2.3 bluebell\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-01-26 16:54:47
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-01-31 20:58:51
 * @FilePath: \2024.1.26 web-app\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"web_app/controller"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	sf "web_app/pkg/snowflake"
	"web_app/routes"
	setting "web_app/settings"

	//"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1 1加载配置
	//1 使用os.Arg提取命令行参数来读取配置文件
	//if len(os.Args) < 2 {
	//	fmt.Println("请指定配置文件路径")
	//}

	//if err := setting.Init(os.Args[1]); err != nil {   使用os.Arg[1]获取到的命令行参数中的filepath

	//使用flag 获取到的filepath
	if err := setting.Init(setting.Setflag()); err != nil {
		fmt.Printf("init setting failed err :%v\n", err)
	}
	//2 初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("Init logger failed ,err :%v\n", err)
	}
	defer zap.L().Sync()

	//3 初始化mysql连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("Init mysql failed ,err :%v\n", err)
	}
	defer mysql.Close() //关闭close

	//雪花算法
	if err := sf.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed ,err:%v\n", err)
	}

	//4 初始化redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("Init redis failed ,err :%v\n", err)
	}
	defer redis.Close()

	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed err:%v\n", err)
	}

	//5注册路由
	r := routes.SetUpRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed err:#{err}\b")
		return
	}

}
