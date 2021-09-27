package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"uapply_go/dao/mysql"
	"uapply_go/dao/redis"
	"uapply_go/logger"
	"uapply_go/router"
	"uapply_go/setting"
)

// @title uapply_go后端接口文档
// @version 1.0
// @host 121.40.193.220
// @BasePath /api/uapply
func main() {
	// 配置路由
	r := router.SetRouter()

	Init()

	port := viper.GetString("app.port")
	r.Run(":" + port)
}

func Init() {
	// 配置信息
	err := setting.Init()
	if err != nil {
		// zap 还没初始化，先用一下原生库
		log.Printf("%+v \n", err)
		// 配置加载失败就不用干了
		return
	}

	// 日志配置
	err = logger.Init(viper.GetString("log.mode"))
	if err != nil {
		log.Printf("%+v \n", err)
	}

	// mysql
	err = mysql.Init()
	if err != nil {
		// 方便调试
		log.Printf("%+v \n", err)
		// 输出到日志，方便上线后查看
		zap.L().Error("mysql init error", zap.Error(err))
	}
	defer mysql.Close()

	// redis
	err = redis.Init()
	if err != nil {
		log.Printf("%+v \n", err)
		zap.L().Error("redis init error", zap.Error(err))
	}
	defer redis.Close()
}
