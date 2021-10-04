package router

import (
	"github.com/gin-gonic/gin"
	"uapply_go/controller/adminController"
	"uapply_go/controller/departmentController"
	"uapply_go/controller/testController"
	"uapply_go/controller/wxController"
	_ "uapply_go/docs"
	"uapply_go/logger"
	"uapply_go/middleware/auth"
	"uapply_go/middleware/cors"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/*
使用了缓存的地方：
redis：
	adminController.Organizations // 查看组织（新增组织或者减少组织的时候记得改缓存）
	departmentController.Login // 部门登录 （部门改账号密码的时候记得删缓存）
*/
const (
	baseUrl = "/api/uapply"
)

func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), cors.Cors()) // 跨域等等

	// 测试及文档
	test := r.Group("/api/uapply")
	{
		test.GET("/ping", testController.Pong) // 测试
		// http://localhost:9090/api/uapply/swagger/index.html
		test.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) // swagger
	}
	// 部门登录
	r.POST(baseUrl+"/login", departmentController.Login)
	// 管理员选项
	admin := r.Group(baseUrl + "/admin")
	admin.Use(auth.JWTAuthMiddleware())
	{
		// 查看组织
		admin.GET("/organizations", adminController.Organizations)
		// 创建组织
		admin.POST("/organization", adminController.Organization)
		// 创建社团
		admin.POST("/department", adminController.Department)
	}
	// 微信小程序选项
	wx := r.Group(baseUrl + "/wx1")
	{
		// 小程序登录
		wx.GET("/login", wxController.Login1)
	}

	return r
}
