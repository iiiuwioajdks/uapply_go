package router

import (
	"github.com/gin-gonic/gin"
	"uapply_go/controller/testController"
	_ "uapply_go/docs"
	"uapply_go/logger"
	"uapply_go/middleware/cors"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), cors.Cors()) // 跨域等等

	baseTest := r.Group("/api/uapply")
	{
		baseTest.GET("/ping", testController.Pong) // 测试
		// http://localhost:9090/api/uapply/swagger/index.html
		baseTest.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) // swagger
	}

	return r
}
