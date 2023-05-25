package main

import (
	"context"
	"flag"
	"fmt"
	"go_frame/dao/mysql"
	"go_frame/dao/redis"
	"go_frame/logger"
	"go_frame/routes"
	"go_frame/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

// Go Web开发较通用的脚手架模板

func main() {
	// 0. 读取命令行参数
	var configPath string
	flag.StringVar(&configPath, "C", "./conf/config.yaml", "The path of config file.")
	flag.Parse() // 解析命令行参数

	// 1. 加载配置
	if err := settings.Init(configPath); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)

		zap.L().Error("init settings failed", zap.Error(err))
		zap.L().Info("If your config file is not in the current directory, " +
			"please use the -C option to specify the configuration file path.")
		return
	}
	zap.L().Debug("settings init success...")

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		zap.L().Error("init logger failed", zap.Error(err))
		return
	}
	zap.L().Debug("logger init success...")
	defer zap.L().Sync() // 把缓冲区的日志追加到文件中

	// 3. 初始化Mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		zap.L().Error("init mysql failed", zap.Error(err))
		return
	}
	zap.L().Debug("mysql init success...")
	defer mysql.Close() // 程序退出关闭数据库连接

	// 4. 初始化Redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		zap.L().Error("init redis failed", zap.Error(err))
		return
	}
	zap.L().Debug("redis init success...")
	defer redis.Close() // 程序退出关闭redis连接

	// 5. 注册路由
	r := routes.SetUp()
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    viper.GetString("app.port"),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err) // 服务启动失败
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
}
