package router

import (
	"github.com/gin-gonic/gin"
	"uapply_go/controller/adminController"
	"uapply_go/controller/departmentController"
	"uapply_go/controller/testController"
	_ "uapply_go/docs"
	"uapply_go/logger"
	"uapply_go/middleware/cors"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	baseUrl = "/api/uapply"
)

func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), cors.Cors()) // 跨域等等

	test := r.Group("/api/uapply")
	{
		test.GET("/ping", testController.Pong) // 测试
		// http://localhost:9090/api/uapply/swagger/index.html
		test.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) // swagger
	}
	// 部门登录
	r.POST(baseUrl+"/login", departmentController.Login)
	// 组织注册，需要管理员权限
	admin := r.Group(baseUrl + "/admin")
	{
		admin.POST("/organizations", adminController.Organizations)
	}

	return r
}
