/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-03 14:58:10
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-03 15:46:04
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
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	//5注册路由
	r := routes.SetUp()

	//6启动服务
	port := setting.Conf.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

	fmt.Println(setting.Conf.Port)
}
