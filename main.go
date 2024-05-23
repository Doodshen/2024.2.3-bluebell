// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bluebell

// @contact.name liwenzhou
// @contact.url

// @host 127.0.0.1:8081
// @BasePath /api/v1
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
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
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
	r := routes.SetUpRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed err:#{err}\b")
		return
	}

}
